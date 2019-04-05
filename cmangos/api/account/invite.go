package account

import (
  "github.com/google/uuid"

  "metagit.org/blizzlike/cmangos-api/modules/database"
)

type InviteInfo struct {
  Token string `json:"token,omitempty"`
}

func AddAccountToInviteToken(token string, id int64) error {
  stmt, err := database.Api.Prepare(
    "UPDATE invitetoken SET account = ? WHERE token = ?;")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(id, token)
  if err != nil {
    return err
  }

  return nil
}

func WriteInviteToken(id int) (string, error) {
  t, _ := uuid.NewRandom()
  token := t.String()
  stmt, err := database.Api.Prepare(
    "INSERT INTO invitetoken (token, friend) VALUES (?, ?);")
  if err != nil {
    return token, err
  }
  defer stmt.Close()

  _, err = stmt.Exec(token, id)
  if err != nil {
    return token, err
  }

  return token, nil
}
