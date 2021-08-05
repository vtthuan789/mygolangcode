package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/config"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/handlers"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Staring application on port %s", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
