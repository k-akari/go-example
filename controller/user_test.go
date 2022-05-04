package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/k-akari/go-example/repository"
)

var (
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/users/", HandleUsers(&repository.FakeUser{}))
	writer = httptest.NewRecorder()
}

func TestShowUser(t *testing.T) {
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

func TestListUser(t *testing.T) {
	request, _ := http.NewRequest("GET", "/users/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestCreateUser(t *testing.T) {
	json := strings.NewReader(`{"name":"username1", "email":"name1@email.com", "password":"password"}`)
	request, _ := http.NewRequest("POST", "/users/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	json := strings.NewReader(`{"name":"updated_name", "email":"updated_name@email.com"}`)
	request, _ := http.NewRequest("PATCH", "/users/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/users/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
