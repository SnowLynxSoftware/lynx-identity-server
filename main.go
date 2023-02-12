package main

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/config"
	"SnowLynxSoftware/lynx-identity-server/pkg/database"
	"SnowLynxSoftware/lynx-identity-server/pkg/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// Load Configuration
	config.AppConfig = config.LoadEnvIntoConfiguration()

	// Connect to the database
	database.DBConnection = database.GetDBConnection()

	// Setup Gin Routers
	r := routers.SetupRouter()

	// Start Listening on 0.0.0.0:(APP_PORT)
	err := r.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		return
	}
}
