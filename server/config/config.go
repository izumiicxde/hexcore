package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL         string `env:"DB_URL"`
	PORT           string `env:"PORT"`
	API_ENDPOINT   string `env:"API_ENDPOINT"`
	JWT_SECRET     string `env:"JWT_SECRET"`
	RESEND_API_KEY string `env:"RESEND_API_KEY"`
}

var Envs Config
var Validator = validator.New(validator.WithRequiredStructEnabled())

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	if err := env.Parse(&Envs); err != nil {
		panic(err)
	}
}
