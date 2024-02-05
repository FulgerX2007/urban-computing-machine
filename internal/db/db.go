package db

import (
	"database/sql"
	"log"
)

func NewConnection(err error) *sql.DB {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}
