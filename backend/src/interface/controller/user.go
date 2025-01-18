package authController

import (
	"backend/src/usecase"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	UserName string `json:"userId"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	input := usecase.LoginInput{
		UserName: request.UserName,
		Password: request.Password,
	}
	
	output, err := usecase.LoginValidation(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
