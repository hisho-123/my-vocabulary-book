package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
)

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

func LoginValidation(userName string, password string) (userId int, token string, err error) {
	userId, hashPassword, err := gateway.GetUser(userName)
	if err != nil {
		return 0, "", err
	}

	if err := domain.CompareHashPassword(password, hashPassword); err != nil {
		return 0, "", err
	}
	
	token, err = domain.CreateToken(userId)
	if err != nil {
		return 0, "", err
	}
	
	return userId, token, nil
}
