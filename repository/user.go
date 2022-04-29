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
