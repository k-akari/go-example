package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/k-akari/go-example/handler"
	"net/http"
	"time"
)

func main() {
	mux := httprouter.New()

	mux.GET("/users/:id", handler.ShowUser)
	mux.PATCH("/users/:id", handler.UpdateUser)
	mux.DELETE("/users/:id", handler.DeleteUser)
	mux.GET("/users", handler.ShowUsers)
	mux.POST("/users", handler.CreateUser)

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
