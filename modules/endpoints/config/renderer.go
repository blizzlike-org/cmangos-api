package config

import (
  "encoding/json"
  "net/http"
)

type JsonConfigResp struct {
  NeedInvite bool `json:"needInvite,omitempty"`
  Realmd string `json:"realmd,omitempty"`
}

var NeedInvite bool
var Realmd string

func DoConfig(w http.ResponseWriter, r *http.Request) {
  var resp JsonConfigResp
  resp.NeedInvite = NeedInvite
  resp.Realmd = Realmd

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}
