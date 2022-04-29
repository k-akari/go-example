package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/k-akari/go-example/repository"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ShowUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println("Cannot find id of user")
		return
	}

	user, err := repository.UserById(id)
	if err != nil {
		fmt.Println("Cannot find user")
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println("Cannot find id of user")
		return
	}

	if err := repository.DeleteUser(id); err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func ShowUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := repository.AllUsers()
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
}

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var reqParams userCreateRequest
	if err := json.Unmarshal(body, &reqParams); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user := repository.User{Name: reqParams.Name, Email: reqParams.Email, Password: reqParams.Password}
	response, err := repository.InsertUser(user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&response); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprint(w, buf.String())
}
