package dbrepo

import (
	"database/sql"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/config"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
