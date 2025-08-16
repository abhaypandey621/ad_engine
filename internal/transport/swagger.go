package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SwaggerHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/adservice.yaml")
	})
	return r
}
