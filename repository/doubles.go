package repository

import "time"

type FakeUser struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *FakeUser) Fetch(id int) (err error) {
	user.ID = id
	return
}

func (user *FakeUser) List() (users []User, err error) {
	users = []User{
		{ID: 1, UUID: "uuid-001", Name: "test_name1", Email: "test_name1@example.co.jp", Password: "password1", CreatedAt: time.Now()},
		{ID: 2, UUID: "uuid-002", Name: "test_name2", Email: "test_name2@example.co.jp", Password: "password2", CreatedAt: time.Now()},
	}
	return
}

func (user *FakeUser) Create() (err error) {
	return
}

func (user *FakeUser) Update() (err error) {
	return
}

func (user *FakeUser) Delete() (err error) {
	return
}
