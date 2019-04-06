package realm

import (
  "fmt"
  "os"
  "time"
  "database/sql"

  "metagit.org/blizzlike/cmangos-api/cmangos"
  "metagit.org/blizzlike/cmangos-api/modules/database"
  "metagit.org/blizzlike/cmangos-api/modules/config"
)

var realmlist []Realm

func FetchRealms() ([]Realm, error) {
  var rl []Realm
  stmt, err := database.Realmd.Prepare(
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
    err = rows.Scan(&realm.Id, &realm.Name, &realm.Host.Address,
      &realm.Host.Port, &realm.Icon, &realm.Population)
    if err != nil {
      return rl, err
    }

    for _, v := range database.Mangosd {
      if v.Id == realm.Id {
        realm.CharacterInstance.Db = v.Character
      }
    }

    cmangos.CheckDaemon(&realm.Host, time.Duration(config.Settings.Api.CheckTimeout))
    rl = append(rl, realm)
  }

  return rl, nil
}

func GetRealms() []Realm {
  return realmlist
}

func PollRealmStates(interval time.Duration) {
  t := time.Duration(1 * time.Second)
  for range time.Tick(t){
    rl, err := FetchRealms()
    if err != nil {
      fmt.Fprintf(os.Stderr, "Cannot fetch realmlist (%v)\n", err)
      continue
    }
    fmt.Fprintf(os.Stdout, "Fetched realmlist\n")

    realmlist = rl
    t = time.Duration(interval) * time.Second
  }
}
