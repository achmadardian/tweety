package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init() {
	env := os.Getenv("ENVIRONMENT")

	if env == "release" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		Log = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	}
}
