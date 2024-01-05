package config

import "github.com/spf13/viper"

type ConfigEnv struct {
	Host          string `mapstructure:"DB_HOST"`
	Port          string `mapstructure:"DB_PORT"`
	User          string `mapstructure:"DB_USER"`
	Password      string `mapstructure:"DB_PASSWORD"`
	DBname        string `mapstructure:"DB_NAME"`
	Issuer        string `mapstructure:"ISSUER"`
	LibSecretKey  string `mapstructure:"SECRET_KEY"`
	Duration      int    `mapstructure:"EXPIRY"`
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
}

var Config = NewConfig()

// func NewConfig() ConfigEnv {
// 	var Config ConfigEnv

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	expiry, err := strconv.Atoi(os.Getenv("EXPIRY"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	Config.Host = os.Getenv("DB_HOST")
// 	Config.Port = os.Getenv("DB_PORT")
// 	Config.User = os.Getenv("DB_USER")
// 	Config.Password = os.Getenv("DB_PASSWORD")
// 	Config.DBname = os.Getenv("DB_NAME")
// 	Config.Issuer = os.Getenv("ISSUER")
// 	Config.LibSecretKey = os.Getenv("SECRET_KEY")
// 	Config.Duration = expiry
// 	Config.AdminEmail = os.Getenv("ADMIN_EMAIL")
// 	Config.AdminPassword = os.Getenv("ADMIN_PASSWORD")

// 	return Config
// }

func NewConfig() *ConfigEnv {
	var conf *ConfigEnv

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}
