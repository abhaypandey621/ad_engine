package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler returns a chi.Router with all endpoints mounted.
func Handler(ep kitendpoint.Endpoint) http.Handler {
	r := chi.NewRouter()

	r.Mount("/v1/delivery", AdHandler(ep))
	r.Mount("/health", HealthHandler())
	r.Mount("/swagger", SwaggerHandler())

	r.Handle("/metrics", promhttp.Handler())

	return r
}
