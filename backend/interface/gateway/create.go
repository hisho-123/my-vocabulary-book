package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

// 単語帳の作成
func CreateBookByUserId(book domain.CreateBookInput) error {
	const MAX_BOOK_NAME int = 20
	if len(book.BookName) > MAX_BOOK_NAME {
		log.Printf("error: Book name too long.")
		return fmt.Errorf(domain.BadRequest)
	}

	db := db.OpenDB()
	defer db.Close()

	queryCreateBook := "insert into books (user_id, book_name) values (?, ?);"

	_, err := db.Exec(queryCreateBook, strconv.Itoa(book.UserId), book.BookName)
	if err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(domain.InternalServerError)
	}

	var bookId string
	queryGetBookId := "select book_id from books where book_name = ?"
	bookIdRow := db.QueryRow(queryGetBookId, book.BookName)
	if err := bookIdRow.Scan(&bookId); err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(domain.InternalServerError)
	}

	queryCreateWord := "insert into words (book_id, word, translated_word) values (?, ?, ?)"
	for _, v := range book.Words {
		_, err := db.Query(queryCreateWord, bookId, v.Word, v.Translated)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				if mysqlErr.Number == 1406 {
					log.Println("error: ", err)
					return fmt.Errorf(domain.BadRequest)
				}

				log.Println("error: ", err)
				return fmt.Errorf(domain.InternalServerError)
			}

			queryDeleteWords := "delete from words where book_id = ?;"
			db.Query(queryDeleteWords, bookId)

			log.Println("error: ", err)
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
