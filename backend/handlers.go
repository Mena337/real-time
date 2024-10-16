package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	FirstName string
	LastName  string
	Nickname  string
	Age       int
	Email     string
	Password  string
	Gender    string
}

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", "./real-forum.db")
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	user := User{
		FirstName: r.FormValue("FirstName"),
		LastName:  r.FormValue("LastName"),
		Nickname:  r.FormValue("nickname"),
		Age:       parseInt(r.FormValue("Age")),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Gender:    r.FormValue("Gender"),
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	// 	return
	// }

	// user.Password = string(hashedPassword)

	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name, nickname, age, email, password, gender) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Nickname, user.Age, user.Email, user.Password, user.Gender)
	if err != nil {
		http.Error(w, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Registration successful!")
}

func parseInt(value string) int {
	age, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return age
}

