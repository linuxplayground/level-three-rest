package main

import (
  "net/http"
  "github.com/husobee/vestigo"
  "github.com/justinas/alice"
  "encoding/json"
  "io/ioutil"
  "time"
  "log"
)

type User struct {
  Id                string    `json:"id"`
  Name              string    `json:"name"`
  CreationTime      time.Time `json:"creationTime"`
  ModificationTime  time.Time `json:"modificationTime"`
}

var users []User

func main() {
  chain := alice.New(loggingHandler, recoverHandler).Then(CreateRouter())
  http.ListenAndServe(":8080", chain)
}

func loggingHandler(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    log.Printf("Started %s %s", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), time.Since(start))
  }
  return http.HandlerFunc(fn);
}

func recoverHandler(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    defer func() {
      if err := recover(); err != nil {
        log.Printf("panic: %+v", err)
        http.Error(w, http.StatusText(500), 500)
      }
    }()
    next.ServeHTTP(w, r)
  }
  return http.HandlerFunc(fn)
}

func CreateRouter() *vestigo.Router {
  router := vestigo.NewRouter()

  router.Post("/", PostNameHandler)
  return router
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
