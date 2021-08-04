package handlers

import (
	"net/http"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
