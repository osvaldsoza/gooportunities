package router

import (
	"github.com/gin-gonic/gin"
	"github.com/osvaldsoza/gooportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/opening", handler.ShowOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
		api.GET("/openings", handler.ListOpeningsHandler)
	}
}
