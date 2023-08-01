package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func InitLogger(level, component string) error {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.LevelFieldName = "log_level"

	levelParsed, err := zerolog.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("parse log level error %s: %w", level, err)
	}
	zerolog.SetGlobalLevel(levelParsed)
	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("component", component).
		Logger()

	zerolog.DefaultContextLogger = &logger
	return nil
}
