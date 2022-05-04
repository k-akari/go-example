package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k-akari/go-example/repository"
)

func TestShowUser(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", Users)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/users/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var user repository.User
	json.Unmarshal(writer.Body.Bytes(), &user)
	if user.ID != 1 {
		t.Error("Cannot retrieve JSON user")
	}
}
