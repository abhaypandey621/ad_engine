package service

import (
	"context"
	"database/sql"
	"sync"

	"github.com/abhaypandey621/targeting-engine/internal/model"
	"github.com/abhaypandey621/targeting-engine/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	cacheRefreshTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "campaign_cache_refresh_total",
			Help: "Total number of campaign cache refreshes.",
		},
	)
	cacheRefreshErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "campaign_cache_refresh_errors_total",
			Help: "Total number of campaign cache refresh errors.",
		},
	)
	cacheRefreshDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "campaign_cache_refresh_duration_seconds",
			Help:    "Histogram of campaign cache refresh durations.",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(cacheRefreshTotal)
	prometheus.MustRegister(cacheRefreshErrors)
	prometheus.MustRegister(cacheRefreshDuration)
}

type Service interface {
	ServeAd(ctx context.Context, req *model.AdRequest) ([]model.Campaign, error)
	StartCampaignRefresher(ctx context.Context)
}

type service struct {
	database *sql.DB
	cfg      *config.Config

	mu           sync.RWMutex
	activeIndex  int32               // 0 or 1, atomically swapped
	campaignsBuf [2][]model.Campaign // double buffer for campaigns
	appIDMapBuf  [2]map[string]int   // double buffer for appIDMap
}

func NewService(db *sql.DB, cfg *config.Config) Service {
	s := &service{database: db, cfg: cfg}
	// Initial load for both buffers
	s.prepareCampaignsBuffer(0)
	s.prepareCampaignsBuffer(1)
	return s
}

func (s *service) ServeAd(ctx context.Context, req *model.AdRequest) ([]model.Campaign, error) {
	req.Normalize()

	campaigns, appIDMap := s.GetActiveCampaigns()
	appID, ok := appIDMap[req.App]
	if !ok {
		return nil, model.ErrInvalidAppIdentifier
	}

	var result []model.Campaign
	for _, c := range campaigns {
		if MatchCampaign(&c, req, appID) {
			result = append(result, c)
		}
	}
	return result, nil
}
