package main

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite" // Новый драйвер SQLite без CGO
)

func main() {
	db, err := sql.Open("sqlite", "images.db") // Измени "sqlite3" на "sqlite"
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("sqlite"); err != nil { // Измени "sqlite3" на "sqlite"
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations complete")
}
