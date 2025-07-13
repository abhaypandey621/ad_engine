package repository

import "database/sql"

func fetchCountryTargeting(db *sql.DB, campaignID string) (map[string]struct{}, map[string]struct{}) {
	include := make(map[string]struct{})
	exclude := make(map[string]struct{})
	rows, err := db.Query(`SELECT co.country_name, t.inclusion_flag FROM campaign_country_targeting t JOIN country co ON t.country_id = co.country_id WHERE t.campaign_id = ?`, campaignID)
	if err != nil {
		return include, exclude
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var flag bool
		if err := rows.Scan(&name, &flag); err == nil {
			if flag {
				include[name] = struct{}{}
			} else {
				exclude[name] = struct{}{}
			}
		}
	}
	return include, exclude
}

func fetchOSTargeting(db *sql.DB, campaignID string) (map[string]struct{}, map[string]struct{}) {
	include := make(map[string]struct{})
	exclude := make(map[string]struct{})
	rows, err := db.Query(`SELECT o.os_name, t.inclusion_flag FROM campaign_os_targeting t JOIN os o ON t.os_id = o.os_id WHERE t.campaign_id = ?`, campaignID)
	if err != nil {
		return include, exclude
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var flag bool
		if err := rows.Scan(&name, &flag); err == nil {
			if flag {
				include[name] = struct{}{}
			} else {
				exclude[name] = struct{}{}
			}
		}
	}
	return include, exclude
}

func fetchAppTargeting(db *sql.DB, campaignID string) (map[int]struct{}, map[int]struct{}) {
	include := make(map[int]struct{})
	exclude := make(map[int]struct{})
	rows, err := db.Query(`SELECT app_id, inclusion_flag FROM campaign_app_id_targeting WHERE campaign_id = ?`, campaignID)
	if err != nil {
		return include, exclude
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var flag bool
		if err := rows.Scan(&id, &flag); err == nil {
			if flag {
				include[id] = struct{}{}
			} else {
				exclude[id] = struct{}{}
			}
		}
	}
	return include, exclude
}
