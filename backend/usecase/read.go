package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
	"fmt"
	"log"
)

func GetBookList(reaquestHeader string) (bookList []domain.GetBookListOutput, err error) {
	claims, err := domain.ValidateToken(reaquestHeader)
	if err != nil {
		return nil, err
	}

	books, err := gateway.GetBookListByUserId(claims.UserId)
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		bookList = append(bookList, domain.GetBookListOutput{
			BookId:   book.Id,
			BookName: book.Name,
		})
	}

	return bookList, nil
}

func GetBook(requestHeader string, bookId int) (book *domain.GetBookOutput, err error) {
	claims, err := domain.ValidateToken(requestHeader)
	if err != nil {
		return nil, err
	}

	book, err = gateway.GetBookByBookId(bookId)
	if err != nil {
		return nil, err
	}

	if book.UserId != claims.UserId {
		log.Printf("error: UserId in token did not match the one in the bookName: %v.", book.BookName)
		return nil, fmt.Errorf(domain.Unauthorized)
	}

	return book, nil
}
