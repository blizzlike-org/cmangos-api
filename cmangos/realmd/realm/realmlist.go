package realm

import (
  "fmt"
  "net"
  "os"
  "time"
  "database/sql"

  "metagit.org/blizzlike/cmangos-api/modules/database"
  "metagit.org/blizzlike/cmangos-api/modules/config"
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
  d := net.Dialer{Timeout: time.Duration(config.Cfg.Cmangos.Timeout) * time.Second}
  c, err := d.Dial("tcp", fmt.Sprintf("%s:%d", r.Address, r.Port))
  r.Lastcheck = int(time.Now().Unix())
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realm %s (%v)\n", r.Name, err)
    r.State = 0
    return err
  }
  defer c.Close()

  r.State = 1
  return nil
}

var realmlist []Realm

func FetchRealms() ([]Realm, error) {
  var rl []Realm
  stmt, err := database.RealmdDB.Prepare(
    `SELECT id, name, address, port, icon, population
     FROM realmlist
     ORDER BY id ASC;`)
  if err != nil {
    return rl, err
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query()
  for rows.Next() {
    var realm Realm
    err = rows.Scan(&realm.Id, &realm.Name, &realm.Address,
      &realm.Port, &realm.Icon, &realm.Population)
    if err != nil {
      return rl, err
    }

    realm.Check()
    rl = append(rl, realm)
  }

  return rl, nil
}

func GetRealms() []Realm {
  return realmlist
}

func PollRealmStates() {
  t := time.Duration(1 * time.Second)
  for range time.Tick(t){
    rl, err := FetchRealms()
    if err != nil {
      fmt.Fprintf(os.Stderr, "Cannot fetch realmlist (%v)\n", err)
      continue
    }
    fmt.Fprintf(os.Stdout, "Fetched realmlist\n")

    realmlist = rl
    t = time.Duration(config.Cfg.Cmangos.Interval) * time.Second
  }
}
