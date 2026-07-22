package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string `env:"PORT" env-default:"8080"`
	Host string `env:"HOST" env-default:"localhost"`
	// ReadTimeout string `env:"READ_TIMEOUT"`
	// WriteTimeout string `env:"WRITE_TIMEOUT"`
}

type DatabaseConfig struct {
	URL      string `env:"URL" env-required:"true"`
	// MaxOpenConns string `env:"MAX_OPEN_CONNS" env-default:"20"`
	// MaxIdleConns string `env:"MAX_IDLE_CONNS" env-default:"4"`
}

// type App struct {
// 	Environment string //development / production
// 	LogLevel string //debug / info / warn
// }

type Config struct {
	Server   ServerConfig   `env-prefix:"SERVER_"`
	Database DatabaseConfig `env-prefix:"DB_"`
}

func LoadConfig() (*Config, error) {
	//загружает .env в переменные окружения (если файл есть
	_ = godotenv.Load()

	var cfg Config

	// cleanenv.ReadEnv() — читает переменные окружения (и из .env, и системные).
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
