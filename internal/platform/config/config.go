package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port      string `env:"ENV" envDefault:"local"`
	Env       string `env:"PORT" envDefault:"8080"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	JWTSecret string `env:"JWT_SIGNING_KEY,required"`
	DBURL     string `env:"DB_DSN,required"`
	RedisAddr string `env:"REDIS_ADDR,required"`
}

func Load() *Config {

	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return &cfg
}
