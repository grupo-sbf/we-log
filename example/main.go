package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	log "github.com/grupo-sbf/we-log"
)

var (
	logFields = []zap.Field{
		zap.String("LEVEL", log.Configuration.LogLevel),
		zap.String("FORMAT", log.Configuration.LogFormat),
	}
)

// Example_defaultConfig
func Example_defaultConfig() {
	printBetweenNewLines("DEFAULT CONFIG: FORMAT (JSON) // LEVEL (DEBUG)")
	logAllLevels()
}

// Example_jsonInfo
func Example_jsonInfo() {
	printBetweenNewLines("FORMAT (JSON) // LEVEL (INFO)")
	logAllLevels()
}

// Example_consoleDebug
func Example_consoleDebug() {
	printBetweenNewLines("FORMAT (CONSOLE) // LEVEL (DEBUG)")
	logAllLevels()
}

func main() {
	// Run examples with default configuration.
	log.Init()
	Example_defaultConfig()

	// Set LOG_LEVEL to `info`.
	// NOTE: We must reset singleton Log instance, so it can
	// be reseted and reconfigured.
	os.Setenv("LOG_LEVEL", "info")
	log.Log = nil
	log.Init()
	Example_jsonInfo()

	// Set LOG_LEVEL to `debug` again and change
	// LOG_FORMAT to `console`, which will enable development
	// mode and show a more friendly output.
	os.Setenv("LOG_FORMAT", "console")
	os.Setenv("LOG_LEVEL", "debug")
	log.Log = nil
	log.Init()
	Example_consoleDebug()
}

// logAllLevels is used by all of the examples to print
// DEBUG, INFO, WARN and ERROR level logs.
func logAllLevels() {
	log.Log.Debug("test debug", logFields...)
	log.Log.Info("test info", logFields...)
	log.Log.Warn("test warn", logFields...)
	log.Log.Error("test error", logFields...)
}

// printBetweenNewLines is a dummy function used only
// to beautify the output of the examples, by separating them.
func printBetweenNewLines(text string) {
	fmt.Printf("\n%s\n\n", text)
}
