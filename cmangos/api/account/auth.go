package account

import (
  "fmt"
  "os"
  "time"
  "database/sql"

  "metagit.org/blizzlike/cmangos-api/modules/config"
  "metagit.org/blizzlike/cmangos-api/modules/database"
)

func InviteTokenAuth(token string) bool {
  var friend int
  stmt, err := database.Api.Prepare(
    "SELECT friend FROM invitetoken WHERE token = ? AND account IS NULL;")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot query invite tokens (%v)\n", err)
    return false
  }
  defer stmt.Close()

  err = stmt.QueryRow(token).Scan(&friend)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Invalid/Missing invite token (%v)\n", err)
    return false
  }

  return true
}

func AuthenticateByToken(token string) (int, error) {
  var owner, expiry int
  stmt, err := database.Api.Prepare(
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
    stmtUpdate, err = database.Api.Prepare(
      "UPDATE authtoken SET expiry = ? WHERE token = ?;")
    if err != nil {
      return 0, err
    }
    defer stmtUpdate.Close()
    _, err = stmtUpdate.Exec(now + int64(config.Settings.Api.AuthTokenExpiry), token)
    if err != nil {
      return 0, err
    }
  } else {
    stmtDelete, err := database.Api.Prepare(
      "DELETE FROM authtoken WHERE token = ?;")
    if err != nil {
      return 0, err
    }
    defer stmtDelete.Close()
    if err != nil {
      _, err = stmtDelete.Exec(token)
      if err != nil {
        return 0, err
      }
    }
  }

  return owner, nil
}

func WriteAuthToken(token string, id int64) error {
  expiry := time.Now().Unix() + int64(config.Settings.Api.AuthTokenExpiry)
  stmt, err := database.Api.Prepare(
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
