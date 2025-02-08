package routes

import (
	"github.com/gin-gonic/gin"
	"project_vk/backend/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/containers", handlers.GetContainers)
	r.POST("/ping", handlers.PingContainer)

	return r
}
