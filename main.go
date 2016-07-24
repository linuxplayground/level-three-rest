package main

import (
  "net/http"
  "github.com/husobee/vestigo"
  "encoding/json"
  "io/ioutil"
  "crypto/tls"
)

type User struct {
  Name string `json:"name"`
}

var users []User

func main() {
  cfg := &tls.Config{
    MinVersion: tls.VersionTLS12,
    CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
  }

  srv := &http.Server {
		Addr:         ":443",
		Handler:      CreateRouter(),
		TLSConfig:    cfg,
	}

  // TODO - Load these dynamically on start up.
  srv.ListenAndServeTLS("server.pem", "server.key")
}

func CreateRouter() *vestigo.Router {
  router := vestigo.NewRouter()

  router.Get("/:name", GetNameHandler)
  router.Post("/", PostNameHandler)
  return router
}

func GetNameHandler(w http.ResponseWriter, r *http.Request)  {
  name := vestigo.Param(r, "name")

  // TODO - Clean all this up
  AddDefaultHeaders(w)
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

  // TODO - Clean all this up
  AddDefaultHeaders(w)
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

// TODO - Put this in to a chain.
func AddDefaultHeaders(w http.ResponseWriter) {
  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}
