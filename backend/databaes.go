package backend

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTables() {
	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT,
        last_name TEXT,
        nickname TEXT,
        age INTEGER,
        email TEXT UNIQUE,
        password TEXT,
        gender TEXT
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
