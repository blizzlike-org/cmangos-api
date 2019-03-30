package account

type JsonInviteResp struct {
  Token string `json:"token"`
}

func WriteInviteToken(token string, id int) error {
  stmt, err := apiDB.Prepare(
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
