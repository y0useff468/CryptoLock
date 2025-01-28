package Routes

import (
	"github.com/gin-gonic/gin"
	"cryptolock/api/Controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users", Controllers.GetUsers)
		api.GET("/users/:id", Controllers.GetUser)
		api.POST("/users", Controllers.CreateUser)
		api.PUT("/users/:id", Controllers.UpdateUser)
		api.DELETE("/users/:id", Controllers.DeleteUser)
	}
}
