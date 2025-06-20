package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "habits.db")
	if err != nil {
		log.Fatalf("Error with opening DB %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error with connecting to DB %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS habits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		goal_days INTEGER,
		created_at DATETIME
	);`

	if _, err := DB.Exec(createTable); err != nil {
		log.Fatalf("Error with creation of DB %v", err)
	}
}
