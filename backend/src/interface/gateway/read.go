package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetBookListByUserId(userId int) (books []domain.Book, err error) {
	db := db.OpenDB()
	defer db.Close()

	queryGetBookList := "select book_id, book_name, first_review from books where user_id = ?"

	var bookId int
	var bookName string
	var firstReview string

	rows, err := db.Query(queryGetBookList, strconv.Itoa(userId))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(domain.InternalServerError)
	}

	for rows.Next() {
		if err := rows.Scan(&bookId, &bookName, &firstReview); err != nil {
			log.Println(err)
			return nil, fmt.Errorf(domain.InternalServerError)
		}

		books = append(books, domain.Book{
			Id:          &bookId,
			UserId:      userId,
			Name:        bookName,
			FirstReview: &firstReview,
		})
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf(domain.InternalServerError)
	}

	return books, nil
}

func GetBookByBookId(bookId int) (bookName string, words []domain.Word, err error) {
	db := db.OpenDB()
	defer db.Close()

	// 単語帳の名前を取得
	queryGetBook := "select book_name from books where book_id = ?"
	err = db.QueryRow(queryGetBook, bookId).Scan(&bookName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return "", nil, fmt.Errorf(domain.InternalServerError)
		}
		log.Println(err)
		return "", nil, fmt.Errorf(domain.InternalServerError)
	}

	// 単語取得
	queryRowWords := "select word_id, word, translated_word from words where book_id = ?"
	wordsRows, err := db.Query(queryRowWords, bookId)
	if err != nil {
		log.Println(err)
		return "", nil, fmt.Errorf(domain.InternalServerError)
	}

	var wordId int
	var word string
	var translated string

	for wordsRows.Next() {
		if err := wordsRows.Scan(&wordId, &word, &translated); err != nil {
			log.Println(err)
			return "", nil, fmt.Errorf("failed to query get words")
		}

		words = append(words, domain.Word{
			Id:         &wordId,
			BookId:     bookId,
			Word:       word,
			Translated: translated,
		})
	}

	if err := wordsRows.Err(); err != nil {
		log.Println(err)
		return "", nil, fmt.Errorf(domain.InternalServerError)
	}

	return bookName, words, nil
}
