package we_log

import (
	"fmt"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is a singleton instance of zap logger, which intents
// to be used in the whole ecosystem.
var Log *zap.Logger

// Definition of all `ldflags` we use today when building our
// go services.
var (
	Version string
	Commit  string
	Date    string
)

// Init must be called before anything runs and only once when this
// package is used in some application. This function will setup all
// logger configuration based on the environment variables.
//
// NOTE: There is a protection to not do anything if the Log instance is not
// nil, meaning that it was already initialized.
func Init() {
	if Log != nil {
		return
	}

	// Load environment variables using viper.
	setup()

	var err error
	var config zap.Config

	// Manage Log format.
	switch Configuration.LogFormat {
	case "json", "JSON":
		config = zap.NewProductionConfig()
		config.EncoderConfig = zapdriver.NewProductionEncoderConfig()
	default:
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig = zapdriver.NewDevelopmentEncoderConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Manage Log level.
	switch Configuration.LogLevel {
	case "debug", "DEBUG":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	// Load `ldflags` to be printed in all logging calls.
	config.InitialFields = map[string]interface{}{
		"version":      Version,
		"commit":       Commit,
		"version_date": Date,
	}

	Log, err = config.Build(zapdriver.WrapCore(
		zapdriver.ReportAllErrors(true),
		zapdriver.ServiceName(Configuration.ServiceName),
	))
	if err != nil {
		panic(fmt.Sprintf("Error initializing log - %+v", err))
	}
}
