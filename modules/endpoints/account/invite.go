package account

type JsonInviteResp struct {
  Token string `json:"token"`
}

func AddAccountToInviteToken(token string, id int) error {
  stmt, err := ApiDB.Prepare(
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

func WriteInviteToken(token string, id int) error {
  stmt, err := ApiDB.Prepare(
    "INSERT INTO invitetoken (token, friend) VALUES (?, ?);")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(token, id)
  if err != nil {
    return err
  }

  return nil
}
