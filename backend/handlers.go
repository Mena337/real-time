package backend

import (
	"fmt"
	"net/http"
	"real-time-forum/backend/structs"
	"strconv"
)

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
	user := structs.User{
		FirstName: r.FormValue("FirstName"),
		LastName:  r.FormValue("LastName"),
		Nickname:  r.FormValue("nickname"),
		Age:       parseInt(r.FormValue("Age")),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Gender:    r.FormValue("Gender"),
	}

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
