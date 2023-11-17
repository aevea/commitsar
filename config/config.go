package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	// CommitsarConfigPath is used an env variable to override the default location of the config file.
	CommitsarConfigPath = "COMMITSAR_CONFIG_PATH"
)

// LoadConfig iterates through possible config paths. No config will be loaded if no files are present.
func LoadConfig() error {
	viper.AutomaticEnv()
	viper.SetConfigName(".commitsar")
	viper.SetConfigType("yaml")

	if viper.IsSet(CommitsarConfigPath) {
		viper.AddConfigPath(viper.GetString(CommitsarConfigPath))
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("config file not found, using defaults")
		} else {
			// Config file was found but another error was produced
			return err
		}
	}
	return nil
}
