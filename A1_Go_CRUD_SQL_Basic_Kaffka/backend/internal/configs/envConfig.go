package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUser     string `env:"DB_USER,required"`
	DBPass     string `env:"DB_PASS,required"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBName     string `env:"DB_NAME,required"`
	DBType     string `env:"DB_TYPE" envDefault:"mysql"`
	ServerPort string `env:"SERVER_PORT" envDefault:"5000"`
}

func Load() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loding env file")
	}

	cfg := &EnvConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	return cfg
}
