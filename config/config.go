package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort            string `mapstructure:"APP_PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBName             string `mapstructure:"DB_NAME"`
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables.")
	}

	err = viper.Unmarshal(&config)
	return
}
