package main

import (
  "net/http"
  "github.com/husobee/vestigo"
  "github.com/justinas/alice"
  "github.com/bhavikkumar/level-three-rest/common"
)

const (
  ServiceName = "level-three-rest"
)

func main() {
  chain := alice.New(common.TraceHandler, common.RequestTimerHandler).Then(createRouter())
  http.ListenAndServe(":8080", chain)
}

func createRouter() *vestigo.Router {
  router := vestigo.NewRouter()
  router.Get("/" + ServiceName, GetHandler)
  return router
}

func GetHandler(w http.ResponseWriter, r *http.Request)  {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Ok!"))
}
