package routers

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	// Setup Default Router
	r := gin.Default()

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, config.AppConfig)
		//c.String(http.StatusOK, "ok")
	})

	return r
}
