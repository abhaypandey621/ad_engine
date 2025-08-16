package model

type Campaign struct {
	CampaignID     string              `json:"campaign_id"`
	Name           string              `json:"name"`
	ImageURL       string              `json:"image_url"`
	CTA            string              `json:"cta"`
	Status         string              `json:"status"`
	CountryInclude map[string]struct{} `json:"-"`
	CountryExclude map[string]struct{} `json:"-"`
	OSInclude      map[string]struct{} `json:"-"`
	OSExclude      map[string]struct{} `json:"-"`
	AppInclude     map[int]struct{}    `json:"-"`
	AppExclude     map[int]struct{}    `json:"-"`
}
