package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() zerolog.Logger {
	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Лог не открывается")
	}
	logger := zerolog.New(file).With().Timestamp().Logger()
	return logger
}
