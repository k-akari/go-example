package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/k-akari/go-example/repository"
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
