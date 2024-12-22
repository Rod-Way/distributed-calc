package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	RestServerPort string `env:"REST_SERVER_PORT" env-default:"localhost:8080" yaml:"REST_SERVER_PORT"`
}

func New(configName string) *Config {
	cfg := &Config{}
	configPath := fmt.Sprintf(".backend/configs/%s", configName)
	if err := readConfig(configPath, cfg); err != nil {
		log.Printf("warning: failed to read config by path: %d", err)
	} else {
		return cfg
	}

	log.Println("message: trying to read environment variables...")

	if err := readEnvVars(cfg); err != nil {
		log.Printf("warning: failed to read config by path: %d", err)
	} else {
		return cfg
	}

	log.Panicf("error: config not found")

	return nil
}

func readConfig(configPath string, cfg *Config) error {
	return cleanenv.ReadConfig(configPath, &cfg)
}

func readEnvVars(cfg *Config) error {
	return cleanenv.ReadEnv(cfg)
}
