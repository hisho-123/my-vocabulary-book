package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
)

func CreateUser(input domain.UserInput) (*domain.AuthOutput, error) {
	hashedPassword, err := domain.PasswordHash(input.Password)
	if err != nil {
		return nil, err
	}
	user := domain.UserInput{
		UserName: input.UserName,
		Password: hashedPassword,
	}

	userId, err := gateway.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := domain.CreateToken(userId)

	return &domain.AuthOutput{
		UserId: userId,
		Token:  token,
	}, nil
}

func LoginValidation(input domain.UserInput) (*domain.AuthOutput, error) {
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

	return &domain.AuthOutput{
		UserId: userId,
		Token:  token,
	}, nil
}
