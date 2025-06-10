package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
	"fmt"
	"log"
)

func DeleteBook(requestHeader string, input domain.BookInput) error {
	claims, err := domain.ValidateToken(requestHeader)
	if err != nil {
		return err
	}

	book, err := gateway.GetBookByBookId(input.BookId)
	if err != nil {
		return err
	}
	if book.UserId != claims.UserId {
		log.Printf("error: UserId in token did not match the one in the bookId: %v.", input.BookId)
		return fmt.Errorf(domain.Unauthorized)
	}

	err = gateway.DeleteBookByBookId(input.BookId)
	if err != nil {
		return err
	}

	return nil
}
