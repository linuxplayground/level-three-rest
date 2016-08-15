package main

import (
  "net/http"
  "github.com/husobee/vestigo"
  "encoding/json"
  "io/ioutil"
)

type User struct {
  Name string `json:"name"`
}

var users []User

func main() {
  http.ListenAndServe(":8080", CreateRouter())
}

func CreateRouter() *vestigo.Router {
  router := vestigo.NewRouter()

  router.Get("/:name", GetNameHandler)
  router.Post("/", PostNameHandler)
  return router
}

func GetNameHandler(w http.ResponseWriter, r *http.Request)  {
  name := vestigo.Param(r, "name")

  if UserExists(name) {
    w.WriteHeader(200)
    w.Write([]byte(name))
  } else {
    w.WriteHeader(404)
  }
}

func PostNameHandler(w http.ResponseWriter, r *http.Request)  {
  var user User
  body, _ := ioutil.ReadAll(r.Body)
  err := json.Unmarshal(body, &user)
  if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return
  }

  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.Header().Add("Location","/" + user.Name)
  if !UserExists(user.Name) {
    AppendUser(user)
    w.WriteHeader(http.StatusCreated)
  } else {
    w.WriteHeader(http.StatusOK)
  }
  json.NewEncoder(w).Encode(user)
}

func AppendUser(user User) {
  users = append(users, user)
}

func UserExists(name string) bool {
  for _, u := range users {
    if(u.Name == name) {
      return true
    }
  }
  return false
}
