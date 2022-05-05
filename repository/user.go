package repository

import (
	"database/sql"
	"time"
)

type IUser interface {
	Fetch(id int) (err error)
	List() (users []User, err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

type User struct {
	DB        *sql.DB
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Fetch(id int) (err error) {
	err = user.DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (user *User) List() (users []User, err error) {
	rows, err := user.DB.Query("SELECT id, uuid, name, email, password, created_at FROM users")
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
	stmt, err := user.DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.UUID, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
	return
}

func (user *User) Update() (err error) {
	_, err = user.DB.Exec("update users set name = $2, email = $3 where id = $1", user.ID, user.Name, user.Email)
	return
}

func (user *User) Delete() (err error) {
	_, err = user.DB.Exec("delete from users where id = $1", user.ID)
	return
}
