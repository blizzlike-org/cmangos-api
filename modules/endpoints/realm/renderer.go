package realm

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"

  cmangos_realm "metagit.org/blizzlike/cmangos-api/cmangos/realm"
  "metagit.org/blizzlike/cmangos-api/cmangos/realmd"
)

type JsonRealmlistResp struct {
  Realmd realmd.Realmd `json:"realmd,omitempty"`
  Realmlist []cmangos_realm.Realm `json:"realmlist"`
}

func DoRealmlist(w http.ResponseWriter, r *http.Request) {
  realmlist, err := cmangos_realm.FetchRealms()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot fetch realmlist (%v)\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  var resp JsonRealmlistResp
  resp.Realmd = realmd.GetRealmd()
  resp.Realmlist = realmlist

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}