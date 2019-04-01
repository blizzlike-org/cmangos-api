package realm

import (
  "encoding/json"
  "net/http"

  "metagit.org/blizzlike/cmangos-api/cmangos/realmd"
  cmangos_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"
)

func DoRealmlist(w http.ResponseWriter, r *http.Request) {
  realmlist := cmangos_realm.GetRealms()

  var resp realmd.Realmlist
  resp.Realmd = realmd.GetRealmd()
  resp.Realmlist = realmlist

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}
