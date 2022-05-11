package routes

import (
	"github.com/DuduNeves/dev-games.git/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		games := main.Group("games")
		{
			games.GET("/:id", controllers.GetAllGames)
			games.GET("/", controllers.GetGame)
			games.POST("/", controllers.CreateGames)
			games.PUT("/", controllers.UpdateGames)
			games.DELETE("/:id", controllers.DeleteGames)

		}
	}

	return router
}
