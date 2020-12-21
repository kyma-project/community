package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	errWriter := zerolog.ConsoleWriter{
		Out: os.Stderr,
	}
	infoWriter := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatLevel: func(i interface{}) string {
			val, ok := i.(string)
			if !ok {
				return ""
			}
			lvl, err := zerolog.ParseLevel(val)
			if err != nil {
				panic(err)
			}

			if lvl == zerolog.ErrorLevel {
				return "ERROR"
			}
			return ""
		},
	}

	multi := zerolog.MultiLevelWriter(errWriter, infoWriter)

	//zerolog.SyslogLevelWriter(multi)


	logger := zerolog.New(multi).With().Timestamp().Logger()
	ctx := context.WithValue(context.TODO(), "request_id", "2137")

	logger = log.Logger
	logMessages(ctx, logger)
}

func logMessages(ctx context.Context, log zerolog.Logger) {
	enchangeLog(ctx, log.Info()).Msg("log with context")
	log.Info().Str("request_id", "adjhdashkjdas").Msgf("Message: %s", "connected")
	log.Error().Msgf("something bad happen: %s", "log error")
}

func enchangeLog(ctx context.Context, logEvent *zerolog.Event) *zerolog.Event {
	val := ctx.Value("request_id")

	requestID, ok := val.(string)
	if !ok {
		return logEvent.Str("request_id", "not found")
	}
	return logEvent.Str("request_id", requestID)
}
