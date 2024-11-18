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

	resp, err := internal.Call(*addr, *difficulty)
	if err != nil {
		log.Err(err).Msg("client error")
		os.Exit(1)
	}

	log.Info().Msgf("response: %s", resp)
}
