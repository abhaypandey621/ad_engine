package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	return r
}
