package handlers

import (
	"database/sql"
	"net/http"
	"real-time-forum/backend"
	"real-time-forum/backend/structs"
)

func RegisterHandler() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			db, err := sql.Open("sqlite3", "./real-forum.db")
			if err != nil {
				http.Error(w, "Database connection error", http.StatusInternalServerError)
				return
			}
			defer db.Close()

			err = backend.InsertUser(db, structs.User{})
			if err != nil {
				http.Error(w, "Failed to register user", http.StatusInternalServerError)
				return
			}
		}
	})
}
