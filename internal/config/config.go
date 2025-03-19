package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     string `envconfig:"PORT" default:"8080"`
	Database struct {
		Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
		Port     string `envconfig:"DATABASE_PORT" default:"5432"`
		User     string `envconfig:"DATABASE_USER" required:"true"`
		Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
		Name     string `envconfig:"DATABASE_NAME" required:"true"`
	} `envconfig:"database"`
}

func Load() (*Config, error) {
	// Загрузить файл .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	cfg := &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
