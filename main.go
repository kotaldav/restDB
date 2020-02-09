package main

import (
  "net/http"
  "log"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

var (

)

func dbInit() *sql.DB {

  db, err := sql.Open("mysql", "")
  if err != nil {
    log.Fatal(err)
  }

  return db
}

func get(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "get called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "not found"}`))
}

func getTableData(w http.ResponseWriter, r *http.Request) {

}

func main() {

  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("", get).Methods(http.MethodGet)
  api.HandleFunc("", notFound)

  api.HandleFunc("/{database}/{table}", getTableData).Methods(http.MethodGet)

  log.Fatal(http.ListenAndServe(":8080", r))
}
