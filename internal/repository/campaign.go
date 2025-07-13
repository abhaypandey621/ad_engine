package repository

import (
	"database/sql"

	"github.com/abhaypandey621/targeting-engine/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

// FetchAllCampaigns returns all active campaigns and their targeting rules from the DB (for caching).
func FetchAllCampaigns(db *sql.DB) ([]model.Campaign, error) {
	query := `SELECT campaign_id, name, image_url, cta, status FROM campaign_detail WHERE status = 'ACTIVE'`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []model.Campaign
	for rows.Next() {
		var c model.Campaign
		if err := rows.Scan(&c.CampaignID, &c.Name, &c.ImageURL, &c.CTA, &c.Status); err != nil {
			return nil, err
		}
		// Fetch targeting rules for this campaign
		c.CountryInclude, c.CountryExclude = fetchCountryTargeting(db, c.CampaignID)
		c.OSInclude, c.OSExclude = fetchOSTargeting(db, c.CampaignID)
		c.AppInclude, c.AppExclude = fetchAppTargeting(db, c.CampaignID)
		campaigns = append(campaigns, c)
	}
	return campaigns, nil
}
