package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// GloGlobalConfig is the global config
type GlobalConfig struct {
	Database databaseConfig
	Server   serverConfig
}

type databaseConfig struct {
	Dialect string
	DSN     string
}

type serverConfig struct {
	Mode string
}

// global configs
var (
	Config GlobalConfig
)

// Load config from file
func Load(file string) error {
	if _, err := toml.DecodeFile(file, &Config); err != nil {
		return err
	}

	return nil
}

// Init the configuration
func Init() {
	mode := os.Getenv("mode")
	configFile := "config.toml"
	if mode != "" {
		configFile = "config." + mode + ".toml"
	}

	err := Load(configFile)

	if err != nil {
		panic(err)
	}
}
