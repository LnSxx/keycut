package app

import (
	"github.com/go-chi/chi"
)

type App struct {
	Mux *chi.Mux
}

func New(
	mux *chi.Mux,
) *App {
	return &App{
		Mux: mux,
	}
}
