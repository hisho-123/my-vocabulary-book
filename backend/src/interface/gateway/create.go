package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"fmt"
	"log"
	"strconv"
)

// 単語帳の作成
func CreateBookByUserId(userId int, bookName string, words []domain.Word) error {
	db := db.OpenDB()
	defer db.Close()

	queryCreateBook := "insert into books (user_id, book_name) values (?, ?);"

	_, err := db.Exec(queryCreateBook, strconv.Itoa(userId), bookName)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(domain.InternalServerError)
	}

	var bookId string
	queryGetBookId := "select book_id from books where book_name = ?"
	bookIdRow := db.QueryRow(queryGetBookId, bookName)
	if err := bookIdRow.Scan(&bookId); err != nil {
		log.Println(err)
		return fmt.Errorf(domain.InternalServerError)
	}

	queryCreateWord := "insert into words (book_id, word, translated_word) values (?, ?, ?)"
	for _, v := range words {
		_, err := db.Query(queryCreateWord, bookId, v.Word, v.Translated)
		if err != nil {
			queryDeleteWords := "delete from words where book_id = ?;"
			db.Query(queryDeleteWords, bookId)

			log.Println(err)
			return fmt.Errorf(domain.InternalServerError)
		}
	}
	return nil
}

// 既存の単語帳に単語を登録
func CreateWordByBookId(bookId int, word string, translated string) error {
	db := db.OpenDB()
	defer db.Close()

	queryCreateWord := "insert into words (book_id, word, translated_word) values (?, ?, ?)"
	_, err := db.Query(queryCreateWord, bookId, word, translated)
	if err != nil {
		return fmt.Errorf(domain.InternalServerError)
	}
	return nil
}
