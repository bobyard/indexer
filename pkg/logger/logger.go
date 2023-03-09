package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}
