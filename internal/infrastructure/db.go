package infrastructure

import (
	"database/sql"
	"log"
)

func NewDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
