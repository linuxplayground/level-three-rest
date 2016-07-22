package main

import (
  "net/http"
  "net/http/httptest"
	"testing"
  "github.com/husobee/vestigo"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "bytes"
)

var (
  router *vestigo.Router
)

func init() {
  router = CreateRouter()
}

func TestUserDoesNotExists(t *testing.T) {
  exists := UserExists("does-not-exist")
  assert.False(t, exists)
}

func TestAppendUser(t *testing.T) {
  user := User{Name:"test-user"}
  AppendUser(user)
  exists := UserExists(user.Name)
  assert.True(t, exists)
}

func TestPostNameHandlerNewUser(t *testing.T) {
  user := User{Name:"test-user-2"}
  u, _ := json.Marshal(user)
  req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(u))
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)
  assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostNameHandlerAlreadyExists(t *testing.T) {
  user := User{Name:"test-user-2"}
  u, _ := json.Marshal(user)
  req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(u))
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)
  assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostNameHandlerInvalidJson(t *testing.T) {
  user := `{ foo": "bar" }`
  req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(user)))
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)
  assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetNameHandler(t *testing.T) {
  user := User{Name:"test-user-3"}
  u, _ := json.Marshal(user)
  req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(u))
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  gReq, _ := http.NewRequest("GET", "/" + user.Name, nil)
  gW := httptest.NewRecorder()
  router.ServeHTTP(gW, gReq)

  assert.Equal(t, http.StatusOK, gW.Code)
}

func TestGetNameHandlerDoesNotExist(t *testing.T) {
  gReq, _ := http.NewRequest("GET", "/foo", nil)
  gW := httptest.NewRecorder()
  router.ServeHTTP(gW, gReq)

  assert.Equal(t, http.StatusNotFound, gW.Code)
}
