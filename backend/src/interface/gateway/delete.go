package gateway

import (
	"backend/src/infra/db"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteBookByBookId(bookId int) {
	db := db.OpenDB()
	defer db.Close()

	queryDeleteBook := "delete from books where book_id = ?;"
	_, err := db.Query(queryDeleteBook, bookId)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteWordByWordId(wordId int) {
	db := db.OpenDB()
	defer db.Close()

	queryDeleteWord := "delete from words where word_id = ?;"
	_, err := db.Query(queryDeleteWord, wordId)
	if err != nil {
		log.Fatal(err)
	}
}
