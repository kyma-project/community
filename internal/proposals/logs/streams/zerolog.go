package main

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func main() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:     os.Stderr,
		NoColor: true,

	}
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	logger = zerolog.New(os.Stderr).Level(zerolog.ErrorLevel).With().Timestamp().Logger()
	ctx := context.WithValue(context.TODO(), "request_id", "2137")

	testZerolog(ctx, logger)
}

func testZerolog(ctx context.Context, log zerolog.Logger) {
	enchangeLog(ctx, log.Info()).Msg("log with context")
	log.Info().Str("request_id", "adjhdashkjdas").Msgf("Message: %s", "connected")
	log.Error().Msgf("something bad happen: %s", "log error")
	log.Log().Msg("No log level")
}

func enchangeLog(ctx context.Context, logEvent *zerolog.Event) *zerolog.Event {
	val := ctx.Value("request_id")

	requestID, ok := val.(string)
	if !ok {
		return logEvent.Str("request_id", "not found")
	}
	return logEvent.Str("request_id", requestID)
}
