package repository

import "github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
