package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {
	filterLevel := zap.ErrorLevel
	timeEncoder := zapcore.TimeEncoderOfLayout(time.RFC822)

	consoleLog := zapcore.Lock(os.Stderr)
	logFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= filterLevel
	})

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = timeEncoder
	encoder := getEncoder(JSON, cfg)

	core := zapcore.NewTee(zapcore.NewCore(encoder, consoleLog, logFilter))

	logger := zap.New(core).Sugar().With("APP_NAME", "MY_AWESOME_APP")

	testZapLogger(logger)
}

func testZapLogger(log *zap.SugaredLogger) {
	log.With("context","a","a").Infof("just normal log with msg: %s", "Hello From Zap")
	log.Errorf("Error msg: %s", "some error occured")
}

type Encoder int

const (
	JSON    = 1
	CONSOLE = 2
)

func getEncoder(encoder Encoder, cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch encoder {
	case JSON:
		return zapcore.NewJSONEncoder(cfg)
	case CONSOLE:
		return zapcore.NewConsoleEncoder(cfg)
	default:
		panic("unknown encoder")
	}
}

func configureTwoStreams(filterLevel zapcore.Level, encoder zapcore.Encoder) zapcore.Core {
	consoleLog := zapcore.Lock(os.Stdout)
	errLog := zapcore.Lock(os.Stderr)

	logFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= filterLevel
	})

	errFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})

	return zapcore.NewTee(zapcore.NewCore(encoder, consoleLog, logFilter),
		zapcore.NewCore(encoder, errLog, errFilter))
}
