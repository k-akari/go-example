package main

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/lib/pq"
)

type User struct {
    ID   		 int
    Password string
}

func main() {
    var Db *sql.DB
    Db, err := sql.Open("postgres", "host=db user=user password=password dbname=database sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    sql := "SELECT id, password FROM users WHERE id=$1;"

    pstatement, err := Db.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

    queryID := 1
    var user User

    err = pstatement.QueryRow(queryID).Scan(&user.ID, &user.Password)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(user.ID, user.Password)
}