package router

import (
	"github.com/NatalNW7/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(server *gin.Engine){
	handler.InitializeHandler()
	v1 := server.Group("/api/v1")
	{
		// v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		// v1.GET("/openings", handler.ListOpeningsHandler)
		// v1.PUT("/opening", handler.UpdateOpeningHandler)
		// v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}