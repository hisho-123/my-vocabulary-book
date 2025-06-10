package controller

import (
	"backend/src/domain"
	"backend/src/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookListHandler(c *gin.Context) {
	requestHeader := c.GetHeader("Token")
	if requestHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Json header.",
		})
		return
	}

	bookList, err := usecase.GetBookList(requestHeader)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not get book list.",
		})
		return
	}

	c.JSON(http.StatusOK, bookList)
}

func GetBookHandler(c *gin.Context) {
	requestHeader := c.GetHeader("Token")
	if requestHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Json header.",
		})
		return
	}

	bookIdStr := c.Query("bookId")
	bookId, err := strconv.Atoi(bookIdStr)
	if err !=nil {
		c.JSON(statusCode(fmt.Errorf(domain.BadRequest)), gin.H{
			"error": "Could not get book.",
		})
	}

	book, err := usecase.GetBook(requestHeader, bookId)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not get book.",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
