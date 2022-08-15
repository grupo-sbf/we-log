package we_log

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// config struct which represents all
// available configurations of our logger.
type config struct {
	LogLevel    string
	LogFormat   string
	ServiceName string
}

// Configuration is the object that holds all
// configuration unmarshalled from environment variables.
var Configuration config

// bindKeys is used to bind the name of the environment
// variables to the field of our configuration object.
func bindKeys() {
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.BindEnv("LogFormat", "LOG_FORMAT")
	viper.BindEnv("ServiceName", "SERVICE_NAME")
}

// setDefaults sets default values for the environment
// variables:
// - LOG_LEVEL: debug
// - LOF_FORMAT: json
func setDefaults() {
	viper.SetDefault("LogLevel", "debug")
	viper.SetDefault("LogFormat", "json")
}

// setup is called when the logger is initialized and it
// startup all the configuration required by viper to unmarshall
// the environment variables to our config object.
func setup() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindKeys()
	setDefaults()

	viper.AutomaticEnv()

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		panic(fmt.Sprintf("Error loading logging configurations - %+v", err))
	}
}
