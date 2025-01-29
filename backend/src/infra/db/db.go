package db

import (
	"database/sql"
	"log"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.17.0.1:3306)/my_vocabulary_book")
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return db
}
