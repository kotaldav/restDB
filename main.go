package main

import (
  "net/http"
  "database/sql"
  "gopkg.in/yaml.v2"
  "os"
  "fmt"
  "encoding/json"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  log "github.com/sirupsen/logrus"
)

var (
  db *sql.DB
)



func dbInit(cfg Configuration) {

  connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

  var err error
  db, err = sql.Open("mysql", connString)
  if err != nil {
    log.Fatal(err)
  }
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

  vars := mux.Vars(r)
  //dbName := vars["database"]
  dbTable := vars["table"]

  rows, err := db.Query("Select * from " + dbTable)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  columns, _ := rows.Columns()
  dataMap := make([]map[string]interface{}, 0)

  for rows.Next() {
    colVals := make([]interface{}, len(columns))
    colPtrs  := make([]interface{}, len(columns))
    for i, _ := range colVals{
      colPtrs[i] = &colVals[i]
    }

    if err := rows.Scan(colPtrs...); err != nil {
      panic(err.Error())
    }

    rowMap := make(map[string]interface{})
    for i, colName := range columns {
      rowMap[colName] = *colPtrs[i].(*interface{})
    }
    dataMap = append(dataMap, rowMap)
  }

  json.NewEncoder(w).Encode(dataMap)
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

  cfg.Database.User = os.Getenv("DB_USER")
  cfg.Database.Pass = os.Getenv("DB_PASS")

  return cfg
}

func main() {

  log.SetFormatter(&log.JSONFormatter{})
  standardFields := log.Fields{
    "hostname": "test",
    "appname": "restDB",
    "session": "testNum",
  }
  log.SetOutput(os.Stdout)
  log.SetLevel(log.InfoLevel)



  log.Print("Loading config")
  cfg := readConfig("config.yml")

  log.Print("Establishing db connection")
  dbInit(cfg)

  log.Print("Starting server")
  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("/", homeLink).Methods(http.MethodGet)
  api.HandleFunc("/{database}/{table}", getTableData).Methods(http.MethodGet)
  log.WithFields(standardFields).Info(http.ListenAndServe(":8080", r))
}
