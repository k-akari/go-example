package repository

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func UserById(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UpdateUser(inUser User) (outUser User, err error) {
	statement := "update users set name = $2, email = $3 where id = $1 returning *"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(inUser.Id, inUser.Name, inUser.Email).
		Scan(&outUser.Id, &outUser.Uuid, &outUser.Name, &outUser.Email, &outUser.Password, &outUser.CreatedAt)
	return
}

func DeleteUser(id int) (err error) {
	statement := "delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return
}

func AllUsers() (users []User, err error) {
	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Cannot get users from database")
		panic(err.Error())
	}

	user := User{}
	for rows.Next() {
		error := rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if error != nil {
			fmt.Println("Cannot scan users")
		} else {
			users = append(users, user)
		}
	}

	return
}

func InsertUser(inUser User) (outUser User, err error) {
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) returning *"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), inUser.Name, inUser.Email, inUser.Password, time.Now()).
		Scan(&outUser.Id, &outUser.Uuid, &outUser.Name, &outUser.Email, &outUser.Password, &outUser.CreatedAt)
	return
}
