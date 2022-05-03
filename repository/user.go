package repository

import (
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func GetUser(id int) (user User, err error) {
	user = User{}
	err = DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func ListUsers() (users []User, err error) {
	rows, err := DB.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return
		} else {
			users = append(users, user)
		}
	}
	return
}

func (user *User) Create() (err error) {
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.UUID, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
	return
}

func (user *User) Update() (err error) {
	_, err = DB.Exec("update users set name = $2, email = $3 where id = $1", user.ID, user.Name, user.Email)
	return
}

func (user *User) Delete() (err error) {
	_, err = DB.Exec("delete from users where id = $1", user.ID)
	return
}
