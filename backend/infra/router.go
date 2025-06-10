package router

import (
	"backend/src/interface/controller"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, /*TODO: deploy時、変更必須箇所*/
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
