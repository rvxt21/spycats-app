package main

import (
	"github.com/rs/zerolog/log"
	"github.com/rvxt21/sca-agency/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
