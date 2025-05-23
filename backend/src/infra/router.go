package router

import (
	"backend/src/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		// 新規登録
		api.POST("/register", controller.RegisterHandler)
		// ログイン
		api.POST("/login", controller.LoginHandler)
		// ユーザー削除
		api.POST("/user-delete", controller.DeleteUserHandler)
		// 単語帳作成・取得
		api.POST("/book", controller.CreateHandler)
		api.GET("/book", controller.GetBookHandler)
		// 単語帳名取得
		api.GET("/book-list", controller.GetBookListHandler)
		// 単語帳削除
		api.POST("/book-delete", controller.DeleteBookHandler)
	}

	return router
}
