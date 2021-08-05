package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/config"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
