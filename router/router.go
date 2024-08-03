package router

import "github.com/gin-gonic/gin"

func Initialize() {
	server := gin.Default()

	initializeRoutes(server)

	server.Run()
}