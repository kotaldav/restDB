package main

import (
  "net/http"
  "database/sql"
  "gopkg.in/yaml.v2"
  "os"
  "fmt"
  "encoding/json"
  "time"
  "strings"
  "reflect"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  log "github.com/sirupsen/logrus"
)

var (
  db *sql.DB
)



func dbInit(cfg Configuration) {

  connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
  log.Info(connString)

  var err error
  db, err = sql.Open("mysql", connString)
  if err != nil {
    log.Fatal(err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
}

// Logging wrapper
func logger(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.WithFields(log.Fields{"time": time.Since(time.Now()), "Method": r.Method, "Uri": r.RequestURI,}).Info()
    h.ServeHTTP(w, r)
  })
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
  cfg.Database.Host = os.Getenv("DB_HOST")

  return cfg
}

func getSqlTypes( rows *sql.Rows) ([]reflect.Type) {
  cols, _ := rows.ColumnTypes()
  types := make([]reflect.Type, len(cols))
  for i, col := range cols {
    typ := col.ScanType()
    switch typ {
    case reflect.TypeOf(sql.RawBytes{}):
      types[i] = reflect.TypeOf("")       //!!!!!!! Ugliest freaking hack on earth, Fix
    default:
      types[i] = typ
    }
  }
  return types
}

func rowsToMap( rows *sql.Rows ) ([]map[string]interface{}) {

  columns, _ := rows.Columns()
  types := getSqlTypes(rows)
  dataMap := make([]map[string]interface{}, 0)

  values := make([]interface{}, len(columns))
  for rows.Next() {
    for i := range values {
      values[i] = reflect.New(types[i]).Interface()
    }
    err := rows.Scan(values...)
    if err != nil {
      log.Fatal(err)
    }
    rowMap := make(map[string]interface{})
    for i, colName := range columns {
      rowMap[colName] = values[i]
    }
    log.Info(rowMap)
    dataMap = append(dataMap, rowMap)
  }

  return dataMap
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


func getDblist(w http.ResponseWriter, r *http.Request) {
  rows , err := db.Query("SHOW DATABASES")
  if err != nil {
    log.Fatal(err)
  }

  dataMap := rowsToMap(rows)
  json.NewEncoder(w).Encode(dataMap)
}

func getDbTables(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  dbName := vars["database"]

  rows, err := db.Query("SHOW TABLES IN " + dbName)
  if err != nil {
    log.Fatal(err)
  }

  dataMap := rowsToMap(rows)
  json.NewEncoder(w).Encode(dataMap)
}

func getTableData(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  dbName := vars["database"]
  dbTable := vars["table"]

  log.Info(dbTable)

  rows, err := db.Query("SELECT * FROM " + dbName + "." + dbTable)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  w.Header().Set("Content-type","application/json")

  dataMap := rowsToMap(rows)
  json.NewEncoder(w).Encode(dataMap)
}

func insTableData(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  dbName  := vars["database"]
  dbTable := vars["table"]

  body := make(map[string]interface{})
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  log.Info(body)

  var columns []string
  var values  []string
  for k, v := range body {
    columns = append(columns, k)
    values = append(values, v.(string))
  }
  query := "INSERT INTO " + dbName + "." + dbTable + " (" + strings.Join(columns[:], ",") + ") VALUES ('" + strings.Join(values[:], ",'") + "')"
  log.Info(query)
  result, err := db.Query(query)
  log.Print(result)
}

func putTableData(w http.ResponseWriter, r *http.Request) {

}

func ptcTableData(w http.ResponseWriter, r *http.Request) {

}

func delTableData(w http.ResponseWriter, r *http.Request) {

}

func main() {

  /*
  srv := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  */

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
  api.Handle("/", logger(http.HandlerFunc(homeLink))).Methods(http.MethodGet)
  api.Handle("/databases",          logger(http.HandlerFunc(getDblist))).Methods(http.MethodGet)
  api.Handle("/{database}/tables",  logger(http.HandlerFunc(getDbTables))).Methods(http.MethodGet)
  api.Handle("/{database}/{table}", logger(http.HandlerFunc(getTableData))).Methods(http.MethodGet)
  api.Handle("/{database}/{table}", logger(http.HandlerFunc(insTableData))).Methods(http.MethodPost)
  api.Handle("/{database}/{table}", logger(http.HandlerFunc(putTableData))).Methods(http.MethodPut)
  api.Handle("/{database}/{table}", logger(http.HandlerFunc(ptcTableData))).Methods(http.MethodPatch)
  api.Handle("/{database}/{table}", logger(http.HandlerFunc(delTableData))).Methods(http.MethodDelete)
  log.WithFields(standardFields).Info(http.ListenAndServe(":8080", r))
}
