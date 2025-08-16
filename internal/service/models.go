package service

// Campaign represents a campaign with its details.
type Campaign struct {
	CampaignID string `json:"campaign_id"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	CTA        string `json:"cta"`
	Status     string `json:"status"`
}

// TargetingRequest represents the incoming request for ad targeting.
type TargetingRequest struct {
	AppID   int    `json:"app_id"`
	OS      string `json:"os"`
	Country string `json:"country"`
}

// TargetingResponse represents the response containing matching campaigns.
type TargetingResponse struct {
	Campaigns []Campaign `json:"campaigns"`
}
