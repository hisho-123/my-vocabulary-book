package router

import (
	"backend/src/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/login", controller.LoginHandler)
	}
	
	return router
}
