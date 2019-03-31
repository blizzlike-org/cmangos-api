package realm

import (
  "fmt"
  "net"
  "os"
  "time"
  "database/sql"

  "metagit.org/blizzlike/cmangos-api/modules/database"
)

type Realm struct {
  Id int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Address string `json:"address,omitempty"`
  Port int `json:"port,omitempty"`
  Icon int `json:"icon,omitempty"`
  Population int `json:"population,omitempty"`
  State int `json:"state"`
  Lastcheck int `json:"lastcheck"`
}

func (r *Realm) Check() error {
  c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", r.Address, r.Port))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realm %s (%v)\n", r.Name, err)
    r.State = 0
    return err
  }
  defer c.Close()

  r.State = 1
  r.Lastcheck = int(time.Now().Unix())
  return nil
}

func FetchRealms() ([]Realm, error) {
  var realmlist []Realm
  stmt, err := database.RealmdDB.Prepare(
    `SELECT id, name, address, port, icon, population
     FROM realmlist
     ORDER BY id ASC;`)
  if err != nil {
    return realmlist, err
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query()
  for rows.Next() {
    var realm Realm
    err = rows.Scan(&realm.Id, &realm.Name, &realm.Address,
      &realm.Port, &realm.Icon, &realm.Population)
    if err != nil {
      return realmlist, err
    }

    realm.Check()
    realmlist = append(realmlist, realm)
  }

  return realmlist, nil
}
