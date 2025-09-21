package app

import (
	"effective_mobile/internal/config"
	"effective_mobile/internal/handler"
	"effective_mobile/internal/migration"
	"effective_mobile/internal/repository"
	"effective_mobile/internal/service"

	"effective_mobile/pkg/db"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type App struct {
	Handler *handler.Handler
	DB      *sqlx.DB
}

func NewApp(config *config.Config) (*App, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=%s", config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.Dbname, config.Postgres.Sslmode)
	db, err := db.NewDBConnect(dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("db start")
	err = migration.RunMigrations(db, "migrations")
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(&repository.RepositoryDeps{DB: db})
	service := service.NewService(&service.ServiceDeps{Repos: repos})
	handler := handler.NewHandler(service)

	return &App{
		Handler: handler,
		DB:      db,
	}, nil
}

func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}
