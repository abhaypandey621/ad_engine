package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhaypandey621/targeting-engine/internal/model"
	"github.com/go-chi/chi/v5"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	adRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ad_requests_total",
			Help: "Total number of ad delivery requests received.",
		},
		[]string{"status"},
	)
	adRequestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "ad_request_duration_seconds",
			Help:    "Histogram of latencies for ad delivery requests.",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(adRequestsTotal)
	prometheus.MustRegister(adRequestDuration)
}

func AdHandler(ep kitendpoint.Endpoint) http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		app := r.URL.Query().Get("app")
		country := r.URL.Query().Get("country")
		os := r.URL.Query().Get("os")

		if app == "" || country == "" || os == "" {
			adRequestsTotal.WithLabelValues("400").Inc()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "missing required parameter(s): app, country, os"})
			adRequestDuration.Observe(time.Since(start).Seconds())
			return
		}

		req := model.AdRequest{
			App:     app,
			OS:      os,
			Country: country,
		}

		resp, err := ep(context.Background(), req)
		if err != nil {
			if err == model.ErrBadRequest || err == model.ErrInvalidAppIdentifier {
				adRequestsTotal.WithLabelValues("400").Inc()
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				adRequestDuration.Observe(time.Since(start).Seconds())
				return
			}
			adRequestsTotal.WithLabelValues("500").Inc()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			adRequestDuration.Observe(time.Since(start).Seconds())
			return
		}
		response, ok := resp.(model.AdResponse)
		if !ok {
			adRequestsTotal.WithLabelValues("500").Inc()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
			adRequestDuration.Observe(time.Since(start).Seconds())
			return
		}
		if len(response.Campaigns) == 0 {
			adRequestsTotal.WithLabelValues("204").Inc()
			w.WriteHeader(http.StatusNoContent)
			adRequestDuration.Observe(time.Since(start).Seconds())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Campaigns)
		adRequestsTotal.WithLabelValues("200").Inc()
		adRequestDuration.Observe(time.Since(start).Seconds())
	})
	return r
}
