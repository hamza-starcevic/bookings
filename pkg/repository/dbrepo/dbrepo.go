package dbrepo

import (
	"database/sql"

	"github.com/hamza-starcevic/bookings/pkg/config"
	"github.com/hamza-starcevic/bookings/pkg/repository"
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
