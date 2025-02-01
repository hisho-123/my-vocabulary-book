package controller

import (
	"backend/src/domain"
	"backend/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	requestHeader := c.GetHeader("Token")
	if requestHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invaild JSON header.",
		})
		return
	}

	var requestBody domain.CreateBookInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body.",
		})
		return
	}
	
	if err := usecase.CreateBook(requestHeader, requestBody); err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not create book.",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
