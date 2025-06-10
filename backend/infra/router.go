package router

import (
	"backend/src/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Vueの静的ファイル配信（ルートにマウント）
	router.Static("/assets", "./frontend/dist/assets")
	router.StaticFile("/", "./frontend/dist/index.html")

	// catch-all でSPAルーティング対応
	router.NoRoute(func(c *gin.Context) {
			c.File("./frontend/dist/index.html")
	})

	api := router.Group("/api")
	{
		// 新規登録
		api.POST("/register", controller.RegisterHandler)
		// ログイン
		api.POST("/login", controller.LoginHandler)
		// ユーザー削除
		api.DELETE("/user-delete", controller.DeleteUserHandler)
		// 単語帳作成・取得
		api.POST("/book", controller.CreateHandler)
		api.GET("/book", controller.GetBookHandler)
		// 単語帳名取得
		api.GET("/book-list", controller.GetBookListHandler)
		// 単語帳削除
		api.DELETE("/book-delete", controller.DeleteBookHandler)
	}

	return router
}
