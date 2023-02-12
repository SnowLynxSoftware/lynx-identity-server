package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	AppConfig *Config
)

type Config struct {
	AppPort               string `json:"app_port"`
	AppKey                string `json:"app_key"`
	DBConnectionString    string `json:"db_connection_string"`
	RedisConnectionString string `json:"redis_connection_string"`
}

func LoadEnvIntoConfiguration() *Config {
	appPort := os.Getenv("APP_PORT")
	appKey := os.Getenv("APP_KEY")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	redisConnectionString := os.Getenv("REDIS_CONNECTION_STRING")

	appConfig := Config{AppPort: appPort, AppKey: appKey[7:len(appKey)], DBConnectionString: dbConnectionString, RedisConnectionString: redisConnectionString}

	return &appConfig
}
