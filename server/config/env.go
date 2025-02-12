package config

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvStruct struct {
	DB_URL   string `env:"DB_URL"`
	NEON_URL string `env:"NEON_URL"`
	PORT     string `env:"PORT"`

	JWT_SECRET string `env:"JWT_SECRET"`
}

var Envs = InitConfig()

func InitConfig() EnvStruct {
	godotenv.Load()
	s, err := env.ParseAs[EnvStruct]()
	if err != nil {
		slog.Error("error parsing envs", "err", err)
		os.Exit(1)
	}
	return s
}
