package gateway

import (
	"backend/src/domain"
	"backend/src/infra/db"
	"errors"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

// user作成
func CreateUser(user domain.UserInput) (userId int, err error) {
	db := db.OpenDB()
	defer db.Close()

	queryCreateUser := "insert into users (user_name, password) values (?, ?);"

	_, err = db.Exec(queryCreateUser, user.UserName, user.Password)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		errors.As(err, &mysqlErr)
		if mysqlErr.Number == 1062 {
			log.Println("error: ", err)
			return 0, fmt.Errorf(domain.Conflict)
		}

		log.Println("error: ", err)
		return 0, fmt.Errorf(domain.InternalServerError)
	}

	userId, _, err = GetUser(user.UserName)
	if err != nil {
		log.Println("error: ", err)
		return 0, fmt.Errorf(domain.InternalServerError)
	}

	return userId, nil
}

// user取得
func GetUser(userName string) (userId int, hashedPassword string, err error) {
	db := db.OpenDB()
	defer db.Close()

	queryGetUser := "select user_id, password from users where user_name = ?;"
	userRows := db.QueryRow(queryGetUser, userName)

	if err := userRows.Scan(&userId, &hashedPassword); err != nil {
		log.Println("error: ", err)
		return 0, "", fmt.Errorf(domain.InternalServerError)
	}

	return userId, hashedPassword, nil
}

func DeleteUserByUserId(userId int) error {
	db := db.OpenDB()
	defer db.Close()

	queryDeleteUser := "delete from users where user_id = ?;"
	res, err := db.Exec(queryDeleteUser, userId)
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
		log.Printf("no user found with user_id %d", userId)
		return fmt.Errorf(domain.InternalServerError)
	}

	return nil
}
