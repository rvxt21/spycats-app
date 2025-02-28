package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/rvxt21/sca-agency/internal/database"
	"github.com/rvxt21/sca-agency/internal/sca-app/handlers"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
	"github.com/rvxt21/sca-agency/internal/sca-app/storage"
)

type App struct {
	ch   *handlers.CatHandler
	mh   *handlers.MissionHandler
	cs   service.Service
	cst  storage.Storage
	ms   storage.MissionStorage
	msvc service.MissionService
}

func New() (*App, error) {

	connStr := os.Getenv("POSTGRES_CONN_STR")
	if connStr == "" {
		log.Fatal().Msg("POSTGRES_CONN_STR not set in .env file")
	}

	a := &App{}
	

	db := database.DB(connStr)
	mst_ := storage.NewMissionStorage(db)

	//Cat
	st, err := storage.New(db)
	if err != nil {
		return nil, err
	}
	a.cst = st

	a.cs = service.New(a.cst)

	a.ch = handlers.New(a.cs)

	//Mission

	a.ms = *mst_
	a.msvc = *service.NewMissionService(a.ms)
	a.mh = handlers.NewMissionHandler(a.msvc)

	return a, nil
}

func (a *App) Run() error {
	log.Info().Msg("server is running")

	router := gin.Default()

	a.ch.RegisterRoutes(router)
	a.mh.RegisterRoutesM(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	return nil
}
