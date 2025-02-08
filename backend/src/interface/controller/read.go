package controller

import (
	"backend/src/domain"
	"backend/src/usecase"
	"net/http"

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

	var requestBody domain.BookInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body.",
		})
		return
	}

	book, err := usecase.GetBook(requestHeader, requestBody.BookId)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not get book.",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
