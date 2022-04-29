package handler

import (
	"encoding/json"
	"fmt"
	"github.com/k-akari/go-example/repository"
	"net/http"
)

func ShowUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := repository.UserById(1)
	if err != nil {
		fmt.Println("Cannot find user")
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}
