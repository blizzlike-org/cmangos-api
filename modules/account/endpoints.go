package account

import (
  "encoding/base64"
  "encoding/json"
  "net/http"
  "fmt"
  "strings"
  "os"

  "github.com/google/uuid"
)

func DoAuth(w http.ResponseWriter, r *http.Request) {
  auth_header := r.Header.Get("Authorization")
  auth := strings.Split(auth_header, " ")

  if !strings.EqualFold(auth[0], "basic") {
    fmt.Fprintf(os.Stderr, "Authentication method not supported\n")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  credentials, err := base64.StdEncoding.DecodeString(auth[1])
  c := strings.Split(string(credentials), ":")

  id, err := BasicAuth(c[0], c[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot authenticate %s (%v)\n", c[0], err)
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  fmt.Fprintf(os.Stdout, "Authenticated %s\n", c[0])

  t, err := uuid.NewRandom()
  token := t.String()
  err = WriteAuthToken(token, id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot write auth token (%v)\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.Header().Add("X-Auth-Token", token)
  w.WriteHeader(http.StatusOK)
  return
}

func DoInvite(w http.ResponseWriter, r *http.Request) {
  auth_header := r.Header.Get("Authorization")
  auth := strings.Split(auth_header, " ")

  if !strings.EqualFold(auth[0], "token") {
    fmt.Fprintf(os.Stderr, "Authentication method not supported\n")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  id, err := TokenAuth(auth[1])
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot authenticate %s (%v)\n", auth[1], err)
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  fmt.Fprintf(os.Stdout, "Authenticated %s\n", auth[1])

  t, err := uuid.NewRandom()
  token := t.String()
  err = WriteInviteToken(token, id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot write invite token (%v)\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  var inv = JsonInviteResp{token}
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(inv)
}
