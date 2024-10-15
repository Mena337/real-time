package main

import (
	"fmt"
	"log"
	"net/http"
	database "real-time-forum/backend"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database.InitDB("real-forum.db")

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("frontend"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	database.CreateTables()
	log.Println("Database setup complete")

	fmt.Println("Starting server on http://localhost:8080")
	// OpenBrowser("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		r.ParseForm()

// 		user := structs.User

// 		db, err := sql.Open("sqlite3", "./forum.db")
// 		if err != nil {
// 			http.Error(w, "Database connection error", http.StatusInternalServerError)
// 			return
// 		}
// 		defer db.Close()

// 		err = backend.InsertUser(db, user)
// 		if err != nil {
// 			http.Error(w, "Failed to register user", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
