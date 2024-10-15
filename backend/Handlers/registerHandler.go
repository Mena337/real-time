package handlers

import (
	"database/sql"
	"net/http"
	"real-time-forum/backend"
	"real-time-forum/backend/structs"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		user := structs.User{}

		db, err := sql.Open("sqlite3", "./real-forum.db")
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		err = backend.InsertUser(db, user)
		if err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}
	}
}
