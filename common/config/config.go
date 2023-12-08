package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func NewConfig() (*ConfigEnv, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &ConfigEnv{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}, nil
}
