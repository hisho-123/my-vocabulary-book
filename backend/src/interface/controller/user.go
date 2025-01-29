package controller

import (
	"backend/src/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var requestBody LoginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON body",
		})
		return
	}

	input := usecase.LoginInput{
		UserName: requestBody.UserName,
		Password: requestBody.Password,
	}

	output, err := usecase.LoginValidation(input)
	if err != nil {
		c.JSON(statusCode(err), gin.H{
			"error": "permission denied",
		})
	}

	jsonOutput, err := json.Marshal(output)
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON body",
		})
		return
	}

	c.JSON(http.StatusOK, jsonOutput)
}
