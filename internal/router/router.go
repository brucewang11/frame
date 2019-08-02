package router

import (
	"github.com/brucewang11/frame/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()


	// 处理未被捕捉的错误
	//router.Use(nice.Recovery(recoveryHandler))
	app := router.Group("/api/v1")
	authController := app.Group("/auth")
	{
		authController.POST("/add", controller.AddAuth)

	}

	return router
}