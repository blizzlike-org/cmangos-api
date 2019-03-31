package config

import (
  "encoding/json"
  "net/http"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

type JsonConfigResp struct {
  NeedInvite bool `json:"needInvite,omitempty"`
  Realmd string `json:"realmd,omitempty"`
}

func DoConfig(w http.ResponseWriter, r *http.Request) {
  var resp JsonConfigResp
  resp.NeedInvite = config.Cfg.NeedInvite
  resp.Realmd = config.Cfg.Cmangos.Realmd

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}
