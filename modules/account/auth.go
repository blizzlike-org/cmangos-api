package account

import (
  "time"
  "database/sql"
)

func BasicAuth(username, password string) (int, error) {
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

func InviteTokenAuth(token string) bool {
  var friend int
  stmt, err := apiDB.Prepare(
    "SELECT friend FROM invitetoken WHERE token = ? AND account IS NULL;")
  if err != nil {
    return false
  }
  defer stmt.Close()

  err = stmt.QueryRow(token).Scan(&friend)
  if err != nil {
    return false
  }

  return true
}

func TokenAuth(token string) (int, error) {
  var owner, expiry int
  stmt, err := apiDB.Prepare(
    "SELECT owner, expiry FROM authtoken WHERE token = ?;")
  if err != nil {
    return 0, err
  }
  defer stmt.Close()

  err = stmt.QueryRow(token).Scan(&owner, &expiry)
  if err != nil {
    return 0, err
  }

  var stmtUpdate *sql.Stmt
  now := time.Now().Unix()
  if int(now) <= expiry {
    stmtUpdate, err = apiDB.Prepare(
      "UPDATE authtoken SET expiry = ? WHERE token = ?;")
    if err != nil {
      return 0, err
    }
    defer stmtUpdate.Close()
    _, err = stmtUpdate.Exec(now + 3600, token)
    if err != nil {
      return 0, err
    }
  } else {
    stmtUpdate, err = apiDB.Prepare(
      "DELETE FROM authtoken WHERE token = ?;")
    if err != nil {
      return 0, err
    }
    defer stmtUpdate.Close()
    if err != nil {
      _, err = stmtUpdate.Exec(token)
      if err != nil {
        return 0, err
      }
    }
  }

  return owner, nil
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
