package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/k-akari/go-example/repository"
)

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		if path.Base(r.URL.Path) == "users" {
			err = listUsers(w, r)
		} else {
			err = showUser(w, r)
		}
	case "POST":
		err = createUser(w, r)
	case "PATCH":
		err = updateUser(w, r)
	case "DELETE":
		err = deleteUser(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showUser(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := repository.GetUser(id)
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

func listUsers(w http.ResponseWriter, r *http.Request) (err error) {
	users, err := repository.ListUsers()
	if err != nil {
		fmt.Println("Cannot find users")
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

func createUser(w http.ResponseWriter, r *http.Request) (err error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var reqParams userCreateRequest
	if err = json.Unmarshal(body, &reqParams); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user := repository.User{Uuid: repository.CreateUUID(), Name: reqParams.Name, Email: reqParams.Email, Password: reqParams.Password}
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
	w.WriteHeader(201)
	fmt.Fprint(w, buf.String())
	return
}

func updateUser(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var reqParams userUpdateRequest
	if err = json.Unmarshal(body, &reqParams); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user := repository.User{Id: id, Name: reqParams.Name, Email: reqParams.Email}
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

func deleteUser(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := repository.GetUser(id)
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
