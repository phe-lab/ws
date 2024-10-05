package log

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	once   sync.Once
	logger *zerolog.Logger
)

func InitLogger(debug bool) *zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			// log.Logger = log.Logger.With().Caller().Logger()
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		logger = &log.Logger
	})

	return logger
}

func GetLogger() *zerolog.Logger {
	return logger
}
