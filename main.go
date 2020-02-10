package main

import (
  "net/http"
  "log"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  "github.com/spf13/viper"
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

func readConfg() (){

  // Load configuration file
  viper.Setconfigname("config.yml")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()
  viper.SetConfigType("yml")

  // TODO: missing configuration
  var dbConfig DbConfiguration
  err := viper.Unmarshal(&dbConfig)
  if err != nil {
    // TODO: Error handling
  }


}

func main() {



  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("/", homeLink).Methods(http.MethodGet)
  api.HandleFunc("/{database}/{table}", getTableData).Methods(http.MethodGet)
  log.Fatal(http.ListenAndServe(":8080", r))
}
