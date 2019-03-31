package config

import (
  "encoding/json"
  "net/http"
)

type JsonConfigResp struct {
  NeedInvite bool `json:"needInvite,omitempty"`
}

var NeedInvite bool

func DoConfig(w http.ResponseWriter, r *http.Request) {
  var resp JsonConfigResp
  resp.NeedInvite = NeedInvite

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(resp)
  return
}
