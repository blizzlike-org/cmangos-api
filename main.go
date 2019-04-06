package main

import (
  "fmt"
  "log"
  "os"
  "time"
  "net/http"
  "github.com/gorilla/mux"

  cmangos_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"

  "metagit.org/blizzlike/cmangos-api/modules/config"
  "metagit.org/blizzlike/cmangos-api/modules/database"
  e_account "metagit.org/blizzlike/cmangos-api/modules/endpoints/account"
  e_config "metagit.org/blizzlike/cmangos-api/modules/endpoints/config"
  e_realm "metagit.org/blizzlike/cmangos-api/modules/endpoints/realm"
  e_character "metagit.org/blizzlike/cmangos-api/modules/endpoints/realm/character"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "USAGE: %s <config>\n", os.Args[0])
    os.Exit(1)
  }
  _, err := config.Read(os.Args[1])
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

  go cmangos_realm.PollRealmStates(
    time.Duration(config.Settings.Api.CheckInterval))

  router := mux.NewRouter()
  router.HandleFunc("/account", e_account.DoGetAccount).Methods("GET")
  router.HandleFunc("/account", e_account.DoCreateAccount).Methods("POST")
  router.HandleFunc("/account/auth", e_account.DoAuth).Methods("POST")
  router.HandleFunc("/account/invite", e_account.DoInvite).Methods("POST")

  router.HandleFunc("/config", e_config.DoConfig).Methods("GET")

  router.HandleFunc("/realm", e_realm.DoRealmlist).Methods("GET")
  router.HandleFunc("/realm/{realm}/characters/{account}",
    e_character.DoCharacterlistByAccount).Methods("GET")

  log.Fatal(http.ListenAndServe(
    fmt.Sprintf("%s:%d", config.Settings.Api.Listen, config.Settings.Api.Port), router))
}
