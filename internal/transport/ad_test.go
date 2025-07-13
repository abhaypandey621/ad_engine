package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setupTestServer should return an http.Handler for the API, using a test DB or in-memory data.
// You must implement this to wire up your actual handler for integration tests.
func setupTestServer() http.Handler {
	// TODO: Implement this to return your actual router/handler with test DB or mocks
	return nil
}

func TestDeliveryAPI(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Single campaign match",
			query:      "/v1/delivery?app=com.abc.xyz&country=germany&os=android",
			wantStatus: 200,
			wantBody:   `[{"cid":"duolingo","img":"https://somelink2","cta":"Install"}]`,
		},
		{
			name:       "Multiple campaign match",
			query:      "/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android",
			wantStatus: 200,
			wantBody:   `[{"cid":"spotify","img":"https://somelink","cta":"Download"},{"cid":"subwaysurfer","img":"https://somelink3","cta":"Play"}]`,
		},
		{
			name:       "Missing app param",
			query:      "/v1/delivery?country=germany&os=android",
			wantStatus: 400,
			wantBody:   `{"error":"missing app param"}`,
		},
	}

	handler := setupTestServer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.query, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}
			if tt.wantStatus == 200 {
				var got, want []map[string]interface{}
				_ = json.Unmarshal(w.Body.Bytes(), &got)
				_ = json.Unmarshal([]byte(tt.wantBody), &want)
				if len(got) != len(want) {
					t.Errorf("got body %s, want %s", w.Body.String(), tt.wantBody)
				}
			} else {
				if w.Body.String() != tt.wantBody {
					t.Errorf("got body %s, want %s", w.Body.String(), tt.wantBody)
				}
			}
		})
	}
}
