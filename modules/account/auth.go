package account

import (
  "time"
)

func Auth(username, password string) (int, error) {
  var id int
  stmt, err := realmdDB.Prepare(
    `SELECT id FROM account
     WHERE UPPER(username) = UPPER(?) AND
     sha_pass_hash = SHA1(CONCAT(UPPER(?), ':', UPPER(?)));`)
  if err != nil {
    return 0, err
  }
  defer stmt.Close()

  err = stmt.QueryRow(username, username, password).Scan(&id)
  if err != nil {
    return 0, err
  }

  return id, nil
}

func WriteAuthToken(token string, id int) error {
  t := time.Now()
  expiry := t.Unix() + 3600
  stmt, err := apiDB.Prepare(
    "INSERT INTO authtoken (token, owner, expiry) VALUES (?, ?, ?);")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(token, id, expiry)
  if err != nil {
    return err
  }

  return nil
}
