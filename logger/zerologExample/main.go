package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func consolePartsOrder() []string {
	return []string{
		zerolog.LevelFieldName,
		zerolog.TimestampFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
}

func main() {
	zerolog.TimeFieldFormat = ""
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Warn().Str("a", "b").Str("c", "d").Msgf("Hello %s", "world")

	// output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339, PartsOrder: consolePartsOrder()}

	output.FormatCaller = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("| %-12s|:", i)
	}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log := zerolog.New(output).With().Caller().Timestamp().Logger()
	log.Error().Msgf("Hello %s", "world")
	log.Error().Str("a", "b").Str("c", "d").Msgf("Hello %s", "world")
	log.Info().Str("a", "b").Str("c", "d").Msgf("Hello %s", "world")
}
