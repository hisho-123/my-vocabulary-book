package gateway

import (
	"backend/src/infra/db"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func UpdateBook() {
	db := db.OpenDB()
	defer db.Close()

	rows, err := db.Query("select user_id, user_name from user;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}

func UpdateWord() {
	db := db.OpenDB()
	defer db.Close()

	rows, err := db.Query("select user_id, user_name from user;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}
