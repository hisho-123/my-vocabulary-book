package gateway

import (
	"backend/src/infra/db"
)

// user作成
func CreateUserToDB(userName string, hashedPassword string) error {
	db := db.OpenDB()
	defer db.Close()

	queryCreateUser := "insert into users (user_name, password) values (?, ?);"

	_, err := db.Exec(queryCreateUser, userName, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

// user取得
func GetUser(userName string) (userId int, hashedPassword string, err error) {
	db := db.OpenDB()
	defer db.Close()
	
	queryGetUser := "select user_id, password from users where user_name = ?;"
	userRows := db.QueryRow(queryGetUser, userName)

	if err := userRows.Scan(&userId, &hashedPassword); err != nil {
		return 0, "", err
	}

	return userId, hashedPassword, nil
}
