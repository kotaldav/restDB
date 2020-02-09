package main

import (
  "fmt"
  "net/http"
  "log"

  "github.com/gorrila/mux"
)

var (

)

func dbInit() {

}

func get(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(htt.StatusOK)
  w.Write([]byte(`{"message": "get called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(htt.StatusOK)
  w.Write([]byte(`{"message": "not found"}`))
}


func main() {

  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("", get).Methods(http.MethodGet)
  api.HandleFunc("", notFound)

  log.Fatal(http.ListenAndServe(":8080", r))
}
