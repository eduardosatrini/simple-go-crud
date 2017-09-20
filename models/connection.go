package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connection with mysql
func Connection(link string) (*sql.DB, error) {
	db, err := sql.Open("mysql", link)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
