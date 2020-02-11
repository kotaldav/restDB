package main

import (
  "net/http"
  "log"
  "database/sql"
  "gopkg.in/yaml.v2"
  "os"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

var (

)



func dbInit(cfg Configuration) (*sql.DB) {

  connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

  db, err := sql.Open("mysql", connString)
  if err != nil {
    log.Fatal(err)
  }

  return db
}

func homeLink(w http.ResponseWriter, r *http.Request) {
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

func readConfig(filename string) Configuration{

  // Load configuration file
  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  var cfg Configuration
  decoder := yaml.NewDecoder(f)
  err = decoder.Decode(&cfg)
  if err != nil {
    log.Fatal(err)
  }

  return cfg
}

func main() {

  cfg := readConfig("config.yml")

  dbInit(cfg)

  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("/", homeLink).Methods(http.MethodGet)
  api.HandleFunc("/{database}/{table}", getTableData).Methods(http.MethodGet)
  log.Fatal(http.ListenAndServe(":8080", r))
}
