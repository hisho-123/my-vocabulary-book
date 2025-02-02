package usecase

import (
	"backend/src/domain"
	"backend/src/interface/gateway"
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
			BookId: book.Id,
			BookName: book.Name,
		})
	}

	return bookList, nil
}
