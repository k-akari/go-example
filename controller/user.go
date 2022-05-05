// Package controller contains handler functions
package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/k-akari/go-example/repository"
)

func HandleUsers(u repository.IUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			if path.Base(r.URL.Path) == "users" {
				err = listUsers(w, r, u)
			} else {
				err = showUser(w, r, u)
			}
		case "POST":
			err = createUser(w, r, u)
		case "PATCH":
			err = updateUser(w, r, u)
		case "DELETE":
			err = deleteUser(w, r, u)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func showUser(w http.ResponseWriter, r *http.Request, user repository.IUser) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = user.Fetch(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
	return
}

func listUsers(w http.ResponseWriter, r *http.Request, user repository.IUser) (err error) {
	users, err := user.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&users); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
	return
}

func createUser(w http.ResponseWriter, r *http.Request, user repository.IUser) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	if err = json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	err = user.Create()
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err = enc.Encode(&user); err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
	return
}

func updateUser(w http.ResponseWriter, r *http.Request, user repository.IUser) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = user.Fetch(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	if err = json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	err = user.Update()
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, user repository.IUser) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = user.Fetch(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = user.Delete(); err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	return
}
