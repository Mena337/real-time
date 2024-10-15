package backend

import (
	"database/sql"
	"real-time-forum/backend/structs"
)

func InsertUser(db *sql.DB, user structs.User) error {
	query := `
	INSERT INTO users (Nickname, Age, Gender, FirstName, LastName, Email, Password)
	VALUES (?, ?, ?, ?, ?, ?, ?)
`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	return err
}
