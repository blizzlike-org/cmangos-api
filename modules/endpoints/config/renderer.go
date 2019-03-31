package config

import (
  "encoding/json"
  "net/http"

  iface_config "metagit.org/blizzlike/cmangos-api/cmangos/interface/config"
  "metagit.org/blizzlike/cmangos-api/modules/config"
)

func DoConfig(w http.ResponseWriter, r *http.Request) {
  var resp iface_config.InterfaceConfig
  resp.NeedInvite = config.Cfg.NeedInvite
  resp.RealmdAddress = config.Cfg.Cmangos.Realmd

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}
