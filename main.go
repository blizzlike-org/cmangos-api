package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "database/sql"
  "github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql"

  "metagit.org/blizzlike/cmangos-api/modules/config"
  e_account "metagit.org/blizzlike/cmangos-api/modules/endpoints/account"
  e_config "metagit.org/blizzlike/cmangos-api/modules/endpoints/config"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "USAGE: %s <config>\n", os.Args[0])
    os.Exit(1)
  }
  cfg, err := config.Read(os.Args[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to read file %v\n", err)
    os.Exit(2)
  }

  var apiDB, realmdDB *sql.DB
  apiDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    cfg.ApiDB.Username,
    cfg.ApiDB.Password,
    cfg.ApiDB.Hostname,
    cfg.ApiDB.Port,
    cfg.ApiDB.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to api database (%v)\n")
    os.Exit(3)
  }
  defer apiDB.Close()

  realmdDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    cfg.RealmdDB.Username,
    cfg.RealmdDB.Password,
    cfg.RealmdDB.Hostname,
    cfg.RealmdDB.Port,
    cfg.RealmdDB.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realmd database (%v)\n")
    os.Exit(3)
  }
  defer realmdDB.Close()
  e_account.ApiDB = apiDB
  e_account.RealmdDB = realmdDB
  e_account.NeedInvite = cfg.NeedInvite
  e_config.NeedInvite = cfg.NeedInvite

  router := mux.NewRouter()
  router.HandleFunc("/account", e_account.DoCreateAccount).Methods("POST")
  router.HandleFunc("/account/auth", e_account.DoAuth).Methods("POST")
  router.HandleFunc("/account/invite", e_account.DoInvite).Methods("POST")

  router.HandleFunc("/config", e_config.DoConfig).Methods("GET")

  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port), router))
}
