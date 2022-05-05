package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/k-akari/go-example/repository"
	. "gopkg.in/check.v1"
)

type UserTestSuite struct {
	mux    *http.ServeMux
	user   *repository.FakeUser
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&UserTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *UserTestSuite) SetUpTest(c *C) {
	s.user = &repository.FakeUser{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/users/", HandleUsers(&repository.FakeUser{}))
	s.writer = httptest.NewRecorder()
}

func (s *UserTestSuite) TestShowUser(c *C) {
	request, _ := http.NewRequest("GET", "/users/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)

	var user repository.User
	json.Unmarshal(s.writer.Body.Bytes(), &user)
	c.Check(user.ID, Equals, 1)
}

func (s *UserTestSuite) TestListUser(c *C) {
	request, _ := http.NewRequest("GET", "/users/", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
}

func (s *UserTestSuite) TestCreateUser(c *C) {
	json := strings.NewReader(`{"name":"username1", "email":"name1@email.com", "password":"password"}`)
	request, _ := http.NewRequest("POST", "/users/", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
}

func (s *UserTestSuite) TestUpdateUser(c *C) {
	json := strings.NewReader(`{"name":"updated_name", "email":"updated_name@email.com"}`)
	request, _ := http.NewRequest("PATCH", "/users/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
}

func (s *UserTestSuite) TestDeleteUser(c *C) {
	request, _ := http.NewRequest("DELETE", "/users/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
}
