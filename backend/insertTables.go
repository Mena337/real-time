package backend

import (
	"database/sql"
	"real-time-forum/backend/structs"
)

func InsertUser(db *sql.DB, user structs.User) error {
	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name, nickname, age, email, password, gender) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Nickname, user.Age, user.Email, user.Password, user.Gender)
	return err
}
