package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Do nothing
	default:
		t.Errorf("The return value type is not a pointer to chi.Mux, but is %T", v)
	}
}
