package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=db user=user password=password dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
