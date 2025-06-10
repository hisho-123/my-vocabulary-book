package controller

import (
	"backend/src/domain"
	"backend/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var requestBody domain.UserInput
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
		return
	}

	c.JSON(http.StatusOK, output)
}

func LoginHandler(c *gin.Context) {
	var requestBody domain.UserInput
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
		return
	}

	c.JSON(http.StatusOK, output)
}

func DeleteUserHandler(c *gin.Context) {
	requestHeader := c.GetHeader("Token")
	if requestHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Json header.",
		})
		return
	}
	
	if err := usecase.DeleteUser(requestHeader); err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "Could not delete user.",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
