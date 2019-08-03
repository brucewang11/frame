package router

import (
	"github.com/brucewang11/frame/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	app := router.Group("/api/v1")
	app.Use(JWTAuth())
	authController := app.Group("/auth")
	{
		authController.POST("/add", controller.AddAuth)
		authController.POST("/update", controller.UpdateAuth)
		authController.POST("/delete", controller.DelAuth)
		authController.GET("/list", controller.ListAuth)
	}

	return router
}



