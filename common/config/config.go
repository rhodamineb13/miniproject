package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBname       string
	Issuer       string
	LibSecretKey string
	Duration     int
}

var Config *ConfigEnv

func NewConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	expiry, err := strconv.Atoi(os.Getenv("EXPIRY"))
	if err != nil {
		log.Fatal(err)
	}

	Config.Host = os.Getenv("DB_HOST")
	Config.Port = os.Getenv("DB_PORT")
	Config.User = os.Getenv("DB_USER")
	Config.Password = os.Getenv("DB_PASSWORD")
	Config.DBname = os.Getenv("DB_NAME")
	Config.Issuer = os.Getenv("ISSUER")
	Config.LibSecretKey = os.Getenv("SECRET_KEY")
	Config.Duration = expiry
}
