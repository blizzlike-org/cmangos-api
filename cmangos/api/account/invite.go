package account

import (
  "database/sql"

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

func GetInviteTokens(id int) ([]InviteInfo, error) {
  var ii []InviteInfo
  stmt, err := database.Api.Prepare(
    "SELECT token FROM invitetoken WHERE account IS NULL AND friend = ?;")
  if err != nil {
    return ii, err
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query(id)
  for rows.Next() {
    var t InviteInfo
    err = rows.Scan(&t.Token)
    if err != nil {
      return ii, err
    }

    ii = append(ii, t)
  }

  return ii, nil
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
