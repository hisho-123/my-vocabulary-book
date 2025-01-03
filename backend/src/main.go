package main

import (
	// "github.com/gin-gonic/gin"
	"backend/src/interface/gateway"
	"fmt"
)

func main() {
/* 	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てる */
	
	books := gateway.ReadBooks(1)
	for i, v := range books {
		fmt.Printf("key: %v, book: %v\n", i, v.Name)
	}
}
