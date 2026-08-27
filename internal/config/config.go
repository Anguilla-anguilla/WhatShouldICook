package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string `env:"PORT" env-default:"8080"`
	Host string `env:"HOST" env-default:"localhost"`
}

type DatabaseConfig struct {
	URL string `env:"URL" env-required:"true"`
}

type JWTConfig struct {
	Secret string `env:"JWT_SECRET" env-required:"true"`
	TTL    int    `env:"JWT_TTL" env-default:"24"`
}

type Config struct {
	Server   ServerConfig   `env-prefix:"SERVER_"`
	Database DatabaseConfig `env-prefix:"DB_"`
	JWT      JWTConfig
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
