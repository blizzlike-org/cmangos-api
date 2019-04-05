package account

import (
  "encoding/json"
  "net/http"
  "fmt"
  "strings"
  "os"

  api_account "metagit.org/blizzlike/cmangos-api/cmangos/api/account"
  cmangos_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
  "metagit.org/blizzlike/cmangos-api/modules/config"
)

func AuthenticateByToken(w http.ResponseWriter, r *http.Request) (int, error) {
  header := r.Header.Get("Authorization")
  auth := strings.Split(header, " ")
  var id int

  if len(auth) != 2 {
    w.WriteHeader(http.StatusBadRequest)
    return id, fmt.Errorf("Invalid/Missing Authorization header")
  }

  if !strings.EqualFold(auth[0], "token") {
    errmsg := fmt.Sprintf("AUthentication method not supported (%s)", auth[1])
    fmt.Fprintf(os.Stderr, "%s\n", errmsg)
    w.WriteHeader(http.StatusBadRequest)
    return id, fmt.Errorf(errmsg)
  }

  id, err := api_account.AuthenticateByToken(auth[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot authenticate %s (%v)\n", auth[1], err)
    w.WriteHeader(http.StatusUnauthorized)
    return id, err
  }

  return id, nil
}

func DoCreateAccount(w http.ResponseWriter, r *http.Request) {
  var token string
  var err error
  if config.Settings.Api.RequireInvite {
    token, err = AuthenticateByInviteToken(w, r)
    if err != nil {
      return
    }
  }

  var account cmangos_account.AccountInfo
  _ = json.NewDecoder(r.Body).Decode(&account)
  ae, err := cmangos_account.CreateAccount(&account)
  if err != nil {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(ae)
    return
  }

  if config.Settings.Api.RequireInvite {
    api_account.AddAccountToInviteToken(token, account.Id)
  }

  account.Password = ""
  account.Repeat = ""

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(account)
  return
}

func DoGetAccount(w http.ResponseWriter, r *http.Request) {
  id, err := AuthenticateByToken(w, r)
  if err != nil {
    return
  }

  a, err := cmangos_account.GetAccountInfo(id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(a)
  return
}
