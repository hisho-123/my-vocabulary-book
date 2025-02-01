package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
)

type AuthInput struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthOutput struct {
	UserId int `json:"userId"`
	Token  string `json:"token"`
}

func CreateUser(input AuthInput) (*AuthOutput, error) {
	hashedPassword, err := domain.PasswordHash(input.Password)
	if err != nil {
		return nil, err
	}
	userId, err := gateway.CreateUser(input.UserName, hashedPassword)
	if err != nil {
		return nil, err
	}

	token, err := domain.CreateToken(userId)

	return &AuthOutput{
		UserId: userId,
		Token:  token,
	}, nil
}

func LoginValidation(input AuthInput) (*AuthOutput, error) {
	userId, hashPassword, err := gateway.GetUser(input.UserName)
	if err != nil {
		return nil, err
	}

	if err := domain.ComparePassword(input.Password, hashPassword); err != nil {
		return nil, err
	}

	token, err := domain.CreateToken(userId)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		UserId: userId,
		Token:  token,
	}, nil
}
