package router

import (
	"backend/src/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", controller.RegisterHandler)
		api.POST("/login", controller.LoginHandler)
		api.POST("/book", controller.CreateHandler)
	}
	
	return router
}
