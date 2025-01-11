package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"log"
	"strconv"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

func GetBookListByUserId(userId int) []domain.Book {
	db := db.OpenDB()
	defer db.Close()

	queryGetBookList := "select book_id, book_name from books where user_id = ?"

	var books []domain.Book
	var bookId int
	var bookName string
	// var RecentReview time.Time

	rows, err := db.Query(queryGetBookList, strconv.Itoa(userId))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&bookId, &bookName); err != nil {
			log.Fatal(err)
		}

		books = append(books, domain.Book{
			Id:     &bookId,
			UserId: userId,
			Name:   bookName,
			// RecentReview: RecentReview,
		})

	}

	return books
}

func GetBookByBookId(bookId int) (bookName string, words []domain.Word) {
	db := db.OpenDB()
	defer db.Close()

	// 単語帳の名前を取得
	queryGetBook := "select book_name from books where book_id = ?"
	bookNameRow := db.QueryRow(queryGetBook, bookId)

	if err := bookNameRow.Scan(&bookName); err != nil {
		log.Fatal(err)
	}

	// 単語取得
	queryReaWords := "select word_id, word, translated_word from words where book_id = ?"
	wordsRows, err := db.Query(queryReaWords, bookId)
	if err != nil {
		log.Fatal(err)
	}

	var wordId int
	var word string
	var translated string

	for wordsRows.Next() {
		if err := wordsRows.Scan(&wordId, &word, &translated); err != nil {
			log.Fatal(err)
		}

		words = append(words, domain.Word{
			Id:         &wordId,
			BookId:     bookId,
			Word:       word,
			Translated: translated,
		})
	}

	return bookName, words
}
