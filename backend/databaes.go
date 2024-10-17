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
	userTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT,
        last_name TEXT,
        nickname TEXT,
        age INTEGER,
        email TEXT UNIQUE,
        password TEXT,
        gender TEXT
    );`

	categoryTable := `
        CREATE TABLE IF NOT EXISTS categories (
            category_id INTEGER PRIMARY KEY AUTOINCREMENT,
            category_name TEXT NOT NULL UNIQUE
        );`

	commentTable := `
        CREATE TABLE IF NOT EXISTS comments (
            comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
            comment TEXT NOT NULL,
            user_id INTEGER,
            post_id INTEGER,
            FOREIGN KEY (user_id) REFERENCES users(uid),
            FOREIGN KEY (post_id) REFERENCES posts(post_id)
        );`

	likeCommentTable := `
        CREATE TABLE IF NOT EXISTS likeComment (
            comment_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (comment_id) REFERENCES comments(comment_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (comment_id, user_id)
        );`

	dislikeCommentTable := `
        CREATE TABLE IF NOT EXISTS dislikeComment (
            comment_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (comment_id) REFERENCES comments(comment_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (comment_id, user_id)
        );`

	likesTable := `
        CREATE TABLE IF NOT EXISTS likes (
            post_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (post_id, user_id)
        );`

	dislikesTable := `
        CREATE TABLE IF NOT EXISTS dislikes (
            post_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (post_id, user_id)
        );`

	sessionTable := `
        CREATE TABLE IF NOT EXISTS sessions (
            session_id INTEGER PRIMARY KEY AUTOINCREMENT,
            session TEXT,
            user_id INTEGER,
            timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(uid)
        );`

	postTable := `
        CREATE TABLE IF NOT EXISTS posts (
            post_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            dislike INTEGER DEFAULT 0,
            like INTEGER DEFAULT 0,
            post_heading TEXT NOT NULL,
            post_data TEXT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(uid),
            FOREIGN KEY (dislike) REFERENCES dislikes(post_id),
            FOREIGN KEY (like) REFERENCES likes(post_id)
        );`

	postCategoriesTable := `
        CREATE TABLE IF NOT EXISTS post_categories (
            post_id INTEGER,
            category_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (category_id) REFERENCES categories(category_id),
            PRIMARY KEY (post_id, category_id)
        );`

	// Execute table creation statements
	ExecuteSQL(db, userTable)
	ExecuteSQL(db, categoryTable)
	ExecuteSQL(db, commentTable)
	ExecuteSQL(db, likeCommentTable)
	ExecuteSQL(db, dislikeCommentTable)
	ExecuteSQL(db, likesTable)
	ExecuteSQL(db, dislikesTable)
	ExecuteSQL(db, sessionTable)
	ExecuteSQL(db, postTable)
	ExecuteSQL(db, postCategoriesTable)
}


func ExecuteSQL(db *sql.DB, sqlStatement string) {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}