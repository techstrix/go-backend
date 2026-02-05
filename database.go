package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initializeDB() {
	var err error
	db, err = sql.Open("sqlite", "./albums.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS albums (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			artist TEXT,
			price REAL
		);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}
