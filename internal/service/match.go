package service

import "github.com/abhaypandey621/targeting-engine/internal/model"

func MatchCampaign(c *model.Campaign, req *model.AdRequest, appID int) bool {
	// Country targeting
	if len(c.CountryInclude) > 0 {
		if _, ok := c.CountryInclude[req.Country]; !ok {
			return false
		}
	}
	if _, ok := c.CountryExclude[req.Country]; ok {
		return false
	}
	// OS targeting
	if len(c.OSInclude) > 0 {
		if _, ok := c.OSInclude[req.OS]; !ok {
			return false
		}
	}
	if _, ok := c.OSExclude[req.OS]; ok {
		return false
	}
	// App targeting
	if len(c.AppInclude) > 0 {
		if _, ok := c.AppInclude[appID]; !ok {
			return false
		}
	}
	if _, ok := c.AppExclude[appID]; ok {
		return false
	}
	return true
}
