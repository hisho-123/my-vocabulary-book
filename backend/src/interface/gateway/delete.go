package gateway

import (
	"backend/src/infra/db"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteBookByBookId(bookId int) error {
	db := db.OpenDB()
	defer db.Close()

	queryDeleteBook := "delete from books where book_id = ?;"
	res, err := db.Exec(queryDeleteBook, bookId)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to delete book")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to check rows affected")
	}
	if rowsAffected == 0 {
		log.Printf("no book found with book_id %d", bookId)
		return fmt.Errorf("no book found with book_id %d", bookId)
	}

	return nil
}

func DeleteWordByWordId(wordId int) error {
	db := db.OpenDB()
	defer db.Close()

	queryDeleteWord := "delete from words where word_id = ?;"
	res, err := db.Exec(queryDeleteWord, wordId)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to delete word")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to check rows affected")
	}
	if rowsAffected == 0 {
		log.Printf("no word found with word_id %d", wordId)
		return fmt.Errorf("no book found with book_id %d", wordId)
	}

	return nil
}
