package main

import (
	"github.com/gin-gonic/gin"

	"backend/src/domain"
	"backend/src/interface/gateway"
	"fmt"
	"strconv"
)

func main() {
		r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{
	   			"message": "pong",
	   		})
	   	})
		r.Run() // 0.0.0.0:8080 でサーバーを立てる

		

	var word []domain.Word
	for i := 0; i < 5; i++ {
		word = append(word, domain.Word{
			Word:       "word" + strconv.Itoa(i),
			Translated: "translated" + strconv.Itoa(i),
		})
	}

	gateway.CreateBookByUserId(1, "book2", word)

	gateway.CreateWordByBookId(1, "word11", "translated11")
	books, err := gateway.GetBookListByUserId(1)
	if err != nil {
		fmt.Print(err.Error())
	}
	for i, v := range books {
		fmt.Printf("key: %v, book: %v\n", i, v.Name)
	}

	bookName, words, err := gateway.GetBookByBookId(1)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%v\n", bookName)
	for i, v := range words {
		fmt.Printf("key: %v, word: %v\n", i, v.Word)
	}

	gateway.UpdateBookNameByBookId(1, "book1")
	gateway.UpdateWordByWordId(1, "wordUpdated", "translatedUpdated")

	fmt.Printf("----------\n")

	books, err = gateway.GetBookListByUserId(1)
		if err != nil {
		fmt.Print(err.Error())
	}
	for i, v := range books {
		fmt.Printf("key: %v, book: %v\n", i, v.Name)
	}

	bookName, words, err = gateway.GetBookByBookId(1)
		if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%v\n", bookName)
	for i, v := range words {
		fmt.Printf("key: %v, word: %v\n", i, v.Word)
	}

	fmt.Printf("------------\n")
	gateway.DeleteWordByWordId(1)

	bookName, words, err = gateway.GetBookByBookId(1)
		if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%v\n", bookName)
	for i, v := range words {
		fmt.Printf("key: %v, word: %v\n", i, v.Word)
	}

	fmt.Printf("------------\n")
	gateway.DeleteBookByBookId(1)

	books, err = gateway.GetBookListByUserId(1)
		if err != nil {
		fmt.Print(err.Error())
	}
	for i, v := range books {
		fmt.Printf("key: %v, book: %v\n", i, v.Name)
	}
}
