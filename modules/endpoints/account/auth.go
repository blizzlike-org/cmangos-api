package account

import (
  "encoding/base64"
  "net/http"
  "fmt"
  "strings"
  "os"

  "github.com/google/uuid"

  api_account "metagit.org/blizzlike/cmangos-api/cmangos/api/account"
  cmangos_account "metagit.org/blizzlike/cmangos-api/cmangos/realmd/account"
)

func Authenticate(w http.ResponseWriter, r *http.Request) (cmangos_account.AccountInfo, error) {
  header := r.Header.Get("Authorization")
  auth := strings.Split(header, " ")
  var a cmangos_account.AccountInfo

  if len(auth) != 2 {
    w.WriteHeader(http.StatusBadRequest)
    return a, fmt.Errorf("Invalid/Missing Authorization header")
  }

  if !strings.EqualFold(auth[0], "basic") {
    errmsg := fmt.Sprintf("Authentication method not supported (%s)", auth[0])
    fmt.Fprintf(os.Stderr, "%s\n", errmsg)
    w.WriteHeader(http.StatusBadRequest)
    return a, fmt.Errorf(errmsg)
  }

  credentials, err := base64.StdEncoding.DecodeString(auth[1])
  c := strings.Split(string(credentials), ":")
  a, err = cmangos_account.Authenticate(c[0], c[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot authenticate %s (%v)\n", c[0], err)
    w.WriteHeader(http.StatusUnauthorized)
    return a, err
  }

  return a, nil
}

func DoAuth(w http.ResponseWriter, r *http.Request) {
  a, err := Authenticate(w, r)
  if err != nil {
    return
  }

  fmt.Fprintf(os.Stdout, "Authenticated %s\n", a.Username)

  t, err := uuid.NewRandom()
  token := t.String()
  err = api_account.WriteAuthToken(token, a.Id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot write auth token (%v)\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.Header().Add("X-Auth-Token", token)
  w.WriteHeader(http.StatusOK)
  return
}
