package database

import (
	"database/sql"
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
}
