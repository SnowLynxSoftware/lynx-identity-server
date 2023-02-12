package routers

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/auth"
	"SnowLynxSoftware/lynx-identity-server/pkg/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	// Setup Default Router
	r := gin.Default()

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Register New User
	r.POST("/auth/register", func(c *gin.Context) {
		var registerDTO users.UserRegisterDTO
		err := c.Bind(&registerDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		userEntity := auth.RegisterNewUserAccount(&registerDTO)
		c.JSON(http.StatusCreated, userEntity)
	})

	return r
}
