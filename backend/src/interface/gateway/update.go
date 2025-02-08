package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateBookNameByBookId(bookId int, bookName string) error {
	db := db.OpenDB()
	defer db.Close()

	queryUpdateBookName := "update books set book_name = ? where book_id = ?;"
	res, err := db.Exec(queryUpdateBookName, bookName, bookId)
	if err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(domain.InternalServerError)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("error: ", err)
		return fmt.Errorf("failed to query get words")
	}
	if rowsAffected == 0 {
		log.Printf("no book found with book_id %d", bookId)
		return fmt.Errorf(domain.InternalServerError)
	}

	return nil
}

func UpdateWordByWordId(wordId int, word string, translated string) error {
	db := db.OpenDB()
	defer db.Close()

	queryUpdateWord := "update words set word = ?, translated_word = ? where word_id = ?;"
	res, err := db.Exec(queryUpdateWord, word, translated, wordId)
	if err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(domain.InternalServerError)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(domain.InternalServerError)
	}
	if rowsAffected == 0 {
		log.Printf("no word found with word_id %d", wordId)
		return fmt.Errorf(domain.InternalServerError)
	}

	return nil
}
