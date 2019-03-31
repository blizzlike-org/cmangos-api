package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/gorilla/mux"

  "metagit.org/blizzlike/cmangos-api/modules/config"
  "metagit.org/blizzlike/cmangos-api/modules/database"
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

  err = database.Open()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(3)
  }
  defer database.Close()

  router := mux.NewRouter()
  router.HandleFunc("/account", e_account.DoCreateAccount).Methods("POST")
  router.HandleFunc("/account/auth", e_account.DoAuth).Methods("POST")
  router.HandleFunc("/account/invite", e_account.DoInvite).Methods("POST")

  router.HandleFunc("/config", e_config.DoConfig).Methods("GET")

  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port), router))
}
