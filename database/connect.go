package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createDb := `
	CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	_, err = db.Exec(createDb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("!! Create table complete !!")
}
