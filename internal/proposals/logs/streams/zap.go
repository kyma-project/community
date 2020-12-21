package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	filterLevel := zap.InfoLevel

	consoleLog := zapcore.Lock(os.Stdout)
	errLog := zapcore.Lock(os.Stderr)

	logFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= filterLevel
	})

	errFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})

	encoder := getEncoder(JSON, zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(zapcore.NewCore(encoder, consoleLog, logFilter),
		zapcore.NewCore(encoder, errLog, errFilter))

	logger := zap.New(core).Sugar()

	logg := logger.With("APP_NAME", "MY_AWESOME_APP")

	testLogger(logg)

}

func testLogger(log *zap.SugaredLogger) {
	log.Infof("just normal log with msg: %s", "Hello From Zap")
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
