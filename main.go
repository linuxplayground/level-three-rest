package main

import (
  "net/http"
  "github.com/husobee/vestigo"
)

func main() {
  router := vestigo.NewRouter()

  router.Get("/hello/:name", GetWelcomeHandler)

  http.ListenAndServe(":8080", router)
}

func GetWelcomeHandler(w http.ResponseWriter, r *http.Request)  {
  name := vestigo.Param(r, "name")
  w.WriteHeader(200)
  w.Write([]byte("Hello " + name))
}
