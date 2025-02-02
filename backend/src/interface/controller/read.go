package controller

import (
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

	bookList, err := usecase.GetBookList(requestHeader);
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not get book list.",
		})
		return
	}

	c.JSON(http.StatusOK, bookList)
}
