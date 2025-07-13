package model

import "strings"

// AdRequest is the endpoint request struct.
type AdRequest struct {
	App     string `json:"app"`
	OS      string `json:"os"`
	Country string `json:"country"`
}

// AdResponse is the endpoint response struct.
type AdResponse struct {
	Campaigns []Campaign `json:"campaigns,omitempty"`
}

func (r *AdRequest) Normalize() {
	r.App = strings.ToLower(r.App)
	r.OS = strings.ToLower(r.OS)
	r.Country = strings.ToLower(r.Country)
}
