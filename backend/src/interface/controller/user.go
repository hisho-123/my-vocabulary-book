package controller

import (
	"backend/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var requestBody usecase.AuthInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body.",
		})
		return
	}

	output, err := usecase.CreateUser(requestBody)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not create user.",
		})
	}

	c.JSON(http.StatusOK, output)
}

func LoginHandler(c *gin.Context) {
	var requestBody usecase.AuthInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body.",
		})
		return
	}

	output, err := usecase.LoginValidation(requestBody)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Permission denied.",
		})
	}

	c.JSON(http.StatusOK, output)
}
