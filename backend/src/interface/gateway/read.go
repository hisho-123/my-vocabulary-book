package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"log"
	"os"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

func ReadBooks(userId int) []domain.Book{
	db := db.OpenDB()
	defer db.Close()

	query, err := os.ReadFile("interface/gateway/queries/readBooks.sql")
	if err != nil {
		log.Fatal(err)
	}

	var books []domain.Book
	var bookId int
	var bookName string
	// var RecentReview time.Time

	rows, err := db.Query(string(query), userId)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&bookId, &bookName); err != nil {
				log.Fatal(err)
			}

			books = append(books, domain.Book{
				Id:           bookId,
				Name:         bookName,
				// RecentReview: RecentReview,
			})
			
	}
	
	return books
}
