package authController

import (
	"backend/src/usecase"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string `json:"userId"`
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
		switch err.Error() {
			case ""
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
