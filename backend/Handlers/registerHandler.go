package handlers

import (
	"database/sql"
	"net/http"
	"real-time-forum/backend"
	"strconv"
	"real-time-forum/backend/structs"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Extract data from form fields
		user := structs.User{
			FirstName: r.FormValue("FirstName"),
			
			LastName:  r.FormValue("LastName"),
			Age:       func() int {
				age, err := strconv.Atoi(r.FormValue("Age"))
				if err != nil {
					http.Error(w, "Invalid age value", http.StatusBadRequest)
					return 0
				}
				return age
			}(),
			Email:     r.FormValue("email"),
			Password:  r.FormValue("password"), // Hash this password before storing it
			Gender:    r.FormValue("Gender"),
		}

		// Connect to the database
		db, err := sql.Open("sqlite3", "./real-forum.db")
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insert user into the database
		err = backend.InsertUser(db, user)
		if err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		// Successful registration
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered successfully"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
