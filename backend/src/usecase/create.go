package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
)

func CreateBook(requestHeader string, input domain.CreateBookInput) error {
	claims, err := domain.ValidateToken(requestHeader)
	if err != nil {
		return err
	}

	if err := gateway.CreateBookByUserId(claims.UserId, input); err != nil {
		return err
	}
	
	return nil
}
