package log

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	once   sync.Once
	Logger *zerolog.Logger
)

func InitLogger(debug bool) *zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		Logger = &log.Logger
	})

	return Logger
}
