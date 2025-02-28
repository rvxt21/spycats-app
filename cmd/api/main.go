package main

import (
	"github.com/rs/zerolog/log"
	"github.com/rvxt21/sca-agency/internal/sca-app/storage"
)

func main() {
	connStr := "host=database user=TemporaryMainuser password=TemporaryPasw dbname=scaApp port=5432 sslmode=disable"

	_, err := storage.New(connStr)
	if err != nil {
		log.Fatal().Msg("failed to connect to database")
	}
}
