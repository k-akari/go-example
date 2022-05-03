package main

import (
	"net/http"
	"time"

	"github.com/k-akari/go-example/handler"
)

func main() {
	http.HandleFunc("/users/", handler.Users)

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
