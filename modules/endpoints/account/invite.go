package account

import (
  "encoding/json"
  "net/http"
  "fmt"
  "strings"
  "os"

  api_account "metagit.org/blizzlike/cmangos-api/cmangos/api/account"
)

func AuthenticateByInviteToken(w http.ResponseWriter, r *http.Request) (string, error) {
  header := r.Header.Get("Authorization")
  auth := strings.Split(header, " ")

  if len(auth) != 2 {
    w.WriteHeader(http.StatusBadRequest)
    return "", fmt.Errorf("Invalid/Missing Authorization header")
  }

  if !strings.EqualFold(auth[0], "token") {
    errmsg := "Authentication method not supported"
    fmt.Fprintf(os.Stderr, "%s\n", errmsg)
    w.WriteHeader(http.StatusBadRequest)
    return "", fmt.Errorf(errmsg)
  }

  if !api_account.InviteTokenAuth(auth[1]) {
    errmsg := fmt.Sprintf("Cannot authenticate invite %s", auth[1])
    fmt.Fprintf(os.Stderr, "%s\n", errmsg)
    w.WriteHeader(http.StatusUnauthorized)
    return "", fmt.Errorf(errmsg)
  }

  return auth[1], nil
}

func DoInvite(w http.ResponseWriter, r *http.Request) {
  id, err := AuthenticateByToken(w, r)
  if err != nil {
    return
  }

  fmt.Fprintf(os.Stdout, "Authenticated id %d\n", id)

  token, err := api_account.WriteInviteToken(id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot write invite token (%v)\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  var inv = api_account.InviteInfo{token}
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(inv)
  return
}
