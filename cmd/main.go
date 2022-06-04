package main

import (
	"net/http"
	"time"

	"github.com/k-akari/go-example/controller"
	"github.com/k-akari/go-example/repository"
)

func main() {
	http.HandleFunc("/users/", controller.HandleUsers(&repository.User{DB: repository.DB}))

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
