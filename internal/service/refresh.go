package service

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	"github.com/abhaypandey621/targeting-engine/internal/model"
	adDB "github.com/abhaypandey621/targeting-engine/internal/repository"
	"github.com/sirupsen/logrus"
)

func (s *service) StartCampaignRefresher(ctx context.Context) {
	interval := s.cfg.CampaignRefreshInterval()
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.swapAndRefresh()
			}
		}
	}()
	// Initial background load for both buffers
	go s.prepareCampaignsBuffer(0)
	go s.prepareCampaignsBuffer(1)
}

// swapAndRefresh swaps the activeIndex and starts preparing the next buffer in the background.
func (s *service) swapAndRefresh() {
	nextIndex := 1 - s.getActiveIndex()
	start := time.Now()
	// Swap active index
	atomic.StoreInt32(&s.activeIndex, int32(nextIndex))
	// Prepare the other buffer in the background
	go func() {
		err := s.prepareCampaignsBuffer(1 - nextIndex)
		cacheRefreshTotal.Inc()
		cacheRefreshDuration.Observe(time.Since(start).Seconds())
		if err != nil {
			cacheRefreshErrors.Inc()
			logrus.WithError(err).Error("campaign cache refresh error")
		}
	}()
}

// getActiveIndex returns the current active index atomically.
func (s *service) getActiveIndex() int {
	return int(atomic.LoadInt32(&s.activeIndex))
}

// prepareCampaignsBuffer loads campaigns and appIDMap into the specified buffer index.
func (s *service) prepareCampaignsBuffer(index int) error {
	campaigns, err := adDB.FetchAllCampaigns(s.database)
	if err != nil {
		return err
	}
	appIDMap, err := adDB.FetchAllAppIDs(s.database)
	if err != nil {
		return err
	}
	for i := range campaigns {
		campaigns[i].CountryInclude = lowerCaseKeys(campaigns[i].CountryInclude)
		campaigns[i].CountryExclude = lowerCaseKeys(campaigns[i].CountryExclude)
		campaigns[i].OSInclude = lowerCaseKeys(campaigns[i].OSInclude)
		campaigns[i].OSExclude = lowerCaseKeys(campaigns[i].OSExclude)
	}

	s.mu.Lock()
	s.campaignsBuf[index] = campaigns
	s.appIDMapBuf[index] = appIDMap
	s.mu.Unlock()
	return nil
}

// GetActiveCampaigns returns the currently active campaigns and appIDMap.
func (s *service) GetActiveCampaigns() ([]model.Campaign, map[string]int) {
	idx := s.getActiveIndex()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.campaignsBuf[idx], s.appIDMapBuf[idx]
}

func lowerCaseKeys(m map[string]struct{}) map[string]struct{} {
	res := make(map[string]struct{}, len(m))
	for k := range m {
		res[strings.ToLower(k)] = struct{}{}
	}
	return res
}
