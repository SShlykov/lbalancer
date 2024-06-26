package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger Logger `yaml:"logger"`
	App    App    `yaml:"app"`
}

type App struct {
	Hosts []string `yaml:"hosts"`
	Port  string   `yaml:"port"`
}

type Logger struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

func Load(configPath string) (cfg *Config, err error) {
	if configPath == "" {
		return nil, fmt.Errorf("CONFIG_PATH not set")
	}

	configPath += "/default.yml"

	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", configPath)
	}

	cfg = &Config{}
	if err = ReadConfig[Config](configPath, cfg); err != nil {
		return nil, fmt.Errorf("cant read config file: %w", err)
	}

	return cfg, nil
}

func ReadConfig[T any](path string, config *T) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			fmt.Println("Error closing file:", closeErr)
		}
	}()

	return yaml.NewDecoder(f).Decode(config)
}
