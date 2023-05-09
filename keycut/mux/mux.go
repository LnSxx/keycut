package mux

import (
	"keycut/keycut/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func NewMux() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc(http.MethodGet, "/", handlers.RootHandler)

	return r
}
