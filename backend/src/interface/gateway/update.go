package gateway

import (
	"backend/src/infra/db"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func UpdateBookNameByBookId(bookId int, bookName string) {
	db := db.OpenDB()
	defer db.Close()

	queryUpdateBookName := "update books set book_name = ? where book_id = ?;"
	_, err := db.Query(queryUpdateBookName, bookName, bookId)
	if err != nil {
		log.Fatal(err)
	}
}
	
func UpdateWordByWordId(wordId int, word string, translated string) {
	db := db.OpenDB()
	defer db.Close()

	queryUpdateWord := "update words set word = ?, translated_word = ? where word_id = ?;"
	_, err := db.Query(queryUpdateWord, word, translated, wordId)
	if err != nil {
		log.Fatal(err)
	}
}
