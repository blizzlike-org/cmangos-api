package account

import (
  "metagit.org/blizzlike/cmangos-api/modules/database"
)

func Authenticate(username, password string) (AccountInfo, error) {
  var a AccountInfo
  stmt, err := database.Realmd.Prepare(
    `SELECT id, username FROM account
     WHERE UPPER(username) = UPPER(?) AND
     sha_pass_hash = SHA1(CONCAT(UPPER(?), ':', UPPER(?)));`)
  if err != nil {
    return a, err
  }
  defer stmt.Close()

  err = stmt.QueryRow(username, username, password).Scan(
    &a.Id, &a.Username)
  if err != nil {
    return a, err
  }

  return a, nil
}
