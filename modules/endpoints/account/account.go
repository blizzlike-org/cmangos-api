package account

import (
  "regexp"
  "fmt"
  "database/sql"
)

type JsonAccountReq struct {
  Username string `json:"username,omitempty"`
  Password string `json:"password,omitempty"`
  Repeat string `json:"repeat,omitempty"`
  Email string `json:"email,omitempty"`
}

type JsonAccountResp struct {
  Username bool `json:"username"`
  Password bool `json:"password"`
  Repeat bool `json:"repeat"`
  Email bool `json:"email"`
}

var ApiDB *sql.DB
var RealmdDB *sql.DB

func AccountExists(username string) bool {
  var id int
  stmt, err := RealmdDB.Prepare(
    "SELECT id FROM account WHERE username = ?;")
  if err != nil {
    return false
  }
  defer stmt.Close()

  err = stmt.QueryRow(username).Scan(&id)
  if err != nil {
    return false
  }

  return true
}

func CreateAccount(account JsonAccountReq, resp *JsonAccountResp) (int, error) {
  if account.Username == "" || len(account.Username) > 16 ||
     AccountExists(account.Username) {
    resp.Username = false
  }
  if account.Password == "" || len(account.Password) > 16 {
    resp.Password = false
  }
  if account.Password != account.Repeat {
    resp.Repeat = false
  }

  re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  if account.Email == "" || !re.MatchString(account.Email) ||
     EmailExists(account.Email) {
    resp.Email = false
  }

  if !resp.Username || !resp.Password || !resp.Repeat || !resp.Email {
    return 0, fmt.Errorf("Cannot create account")
  }

  stmt, err := RealmdDB.Prepare(
    `INSERT INTO account
     (username, sha_pass_hash, email, joindate)
     VALUES (UPPER(?), SHA1(CONCAT(UPPER(?), ':', UPPER(?))), LOWER(?), NOW());`)
  if err != nil {
    return 0, err
  }
  defer stmt.Close()

  var res sql.Result
  res, err = stmt.Exec(
    account.Username, account.Username, account.Password, account.Email)
  if err != nil {
    return 0, err
  }

  id, _ := res.LastInsertId()
  return int(id), nil
}

func EmailExists(email string) bool {
  var id int
  stmt, err := RealmdDB.Prepare(
    "SELECT id FROM account WHERE email = ?;")
  if err != nil {
    return false
  }
  defer stmt.Close()

  err = stmt.QueryRow(email).Scan(&id)
  if err != nil {
    return false
  }

  return true
}
