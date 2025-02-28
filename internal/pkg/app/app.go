package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/rvxt21/sca-agency/internal/sca-app/handlers"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
	"github.com/rvxt21/sca-agency/internal/sca-app/storage"
)

type App struct {
	h  *handlers.Handler
	s  service.Service
	st storage.Storage
}

func New() (*App, error) {

	connStr := os.Getenv("POSTGRES_CONN_STR")
	if connStr == "" {
		log.Fatal().Msg("POSTGRES_CONN_STR not set in .env file")
	}

	a := &App{}
	st, err := storage.New(connStr)
	if err != nil {
		return nil, err
	}

	a.st = st

	a.s = service.New(a.st)

	a.h = handlers.New(a.s)

	return a, nil
}

func (a *App) Run() error {
	log.Info().Msg("server is running")

	router := gin.Default()

	a.h.RegisterRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	return nil
}
