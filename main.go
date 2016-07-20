package main

import (
  "net/http"
  "github.com/husobee/vestigo"
)

var names []string

func main() {
  router := vestigo.NewRouter()

  router.Get("/:name", GetNameHandler)
  router.Post("/:name", PostNameHandler)

  http.ListenAndServe(":8080", router)
}

func GetNameHandler(w http.ResponseWriter, r *http.Request)  {
  name := vestigo.Param(r, "name")

  if(NameExists(name)) {
    w.WriteHeader(200)
    w.Write([]byte(name))
  } else {
    w.WriteHeader(404)
  }
}

func PostNameHandler(w http.ResponseWriter, r *http.Request)  {
  name := vestigo.Param(r, "name")

  exists := NameExists(name)
  if(!exists) {
    names = append(names, name)
    w.Header().Add("Location","/" + name)
    w.WriteHeader(201)
  } else {
    w.WriteHeader(200)
  }
  w.Write([]byte(name))

}

func NameExists(name string) bool {
  for _, n := range names {
    if(n == name) {
      return true
    }
  }
  return false
}
