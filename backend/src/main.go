package main

import (
	// "github.com/gin-gonic/gin"

	"backend/src/domain"
	"backend/src/interface/gateway"
	"fmt"
	"strconv"
)

func main() {
	/* 	r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{
	   			"message": "pong",
	   		})
	   	})
	   	r.Run() // 0.0.0.0:8080 でサーバーを立てる */

	var word []domain.Word
	for i := 0; i < 5; i++ {
		word = append(word, domain.Word{
			Word:       "word" + strconv.Itoa(i),
			Translated: "translated" + strconv.Itoa(i),
		})
	}

	gateway.CreateBookByUserId(1, "book2", word)

	gateway.CreateWordByBookId(1, "word11", "translated11")

	books := gateway.GetBookListByUserId(1)
	for i, v := range books {
		fmt.Printf("key: %v, book: %v\n", i, v.Name)
	}

	bookName, words := gateway.GetBookByBookId(1)
	fmt.Printf("%v\n", bookName)
	for i, v := range words {
		fmt.Printf("key: %v, word: %v\n", i, v.Word)
	}
}
