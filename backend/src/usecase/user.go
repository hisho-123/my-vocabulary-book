package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
)

type LoginInput struct {
	UserName		string	`json:"userName"`
	Password 	string		`json:"password"`
}

type LoginOutput struct {
	UserId int		`json:"userId"`
	Token string 	`json:"token"`
}

func CreateUser(userName string, password string) error {
	hashedPassword, err := domain.PasswordHash(password)	
	if err != nil {
		return err
	}

	if err := gateway.CreateUserToDB(userName, hashedPassword); err != nil {
		return err
	}
	
	return nil
}

func LoginValidation(input LoginInput) (*LoginOutput, error) {
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
	
	return &LoginOutput{
		UserId: userId,
		Token: token,
	}, nil
}
