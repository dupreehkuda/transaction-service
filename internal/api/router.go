package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a api) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/balance", func(r chi.Router) {
		r.Post("/", a.handlers.FundsHandler)
	})

	return r
}
