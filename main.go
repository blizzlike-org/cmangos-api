package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "database/sql"
  "github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql"

  "metagit.org/blizzlike/cmangos-api/modules/account"

  ini "gopkg.in/ini.v1"
)

var cfg *ini.File

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "USAGE: %s <config>\n", os.Args[0])
    os.Exit(1)
  }
  var err error
  cfg, err = ini.Load(os.Args[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to read file %v\n", err)
    os.Exit(2)
  }

  host := cfg.Section("server").Key("listen").MustString("127.0.0.1")
  port := cfg.Section("server").Key("port").MustInt(5556)

  var apiDB, realmdDB *sql.DB
  apiDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    cfg.Section("apidb").Key("username").MustString("wow-api"),
    cfg.Section("apidb").Key("password").MustString("wow-api"),
    cfg.Section("apidb").Key("hostname").MustString("localhost"),
    cfg.Section("apidb").Key("port").MustInt(3306),
    cfg.Section("apidb").Key("database").MustString("wow-api")))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to api database (%v)\n")
    os.Exit(3)
  }
  defer apiDB.Close()

  realmdDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    cfg.Section("realmddb").Key("username").MustString("mangos"),
    cfg.Section("realmddb").Key("password").MustString("mangos"),
    cfg.Section("realmddb").Key("hostname").MustString("localhost"),
    cfg.Section("realmddb").Key("port").MustInt(3306),
    cfg.Section("realmddb").Key("database").MustString("realmd")))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realmd database (%v)\n")
    os.Exit(3)
  }
  defer realmdDB.Close()

  account.Init(apiDB, realmdDB)

  router := mux.NewRouter()
  router.HandleFunc("/account/auth", account.DoAuth).Methods("POST")
  router.HandleFunc("/account/invite", account.DoInvite).Methods("POST")

  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router))
}
