package main

import (
	"fmt"
	"log"
	"net/http"
	database "real-time-forum/backend"
	"real-time-forum/backend/Handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	database.InitDB("real-forum.db")
	database.CreateTables()
	log.Println("Database setup complete")

	// Serve static files
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("frontend"))))
	http.Handle("/backend/", http.StripPrefix("/backend/", http.FileServer(http.Dir("backend"))))

	// Handle the main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	// Add the registration endpoint
	http.HandleFunc("/register", Handlers.RegisterHandler)

	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
