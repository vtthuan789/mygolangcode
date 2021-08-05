package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/config"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
