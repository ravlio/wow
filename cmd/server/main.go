package main

import (
	"flag"
	"github.com/ravlio/wow/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	addr := flag.String("addr", "", "client address")
	difficulty := flag.Int("difficulty", 4, "difficulty")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	srv := &internal.Server{
		Addr:       *addr,
		Difficulty: *difficulty,
	}

	log.Info().Msgf("server is running on %s with difficulty %d", *addr, *difficulty)
	if err := srv.ListenAndServe(); err != nil {
		log.Err(err).Msg("server error")
		os.Exit(1)
	}
}
