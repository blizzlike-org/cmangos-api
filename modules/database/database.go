package database

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

type RealmDB struct {
  Name string
  Character *sql.DB
  World *sql.DB
}

var ApiDB *sql.DB
var RealmdDB *sql.DB
var RealmsDB []RealmDB

func Close() {
  ApiDB.Close()
  RealmdDB.Close()

  for _, v := range RealmsDB {
    v.Character.Close()
    v.World.Close()
  }
}

func Open() error {
  var err error
  ApiDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    config.Cfg.ApiDB.Username,
    config.Cfg.ApiDB.Password,
    config.Cfg.ApiDB.Hostname,
    config.Cfg.ApiDB.Port,
    config.Cfg.ApiDB.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to api database (%v)\n", err)
    return err
  }

  RealmdDB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    config.Cfg.RealmdDB.Username,
    config.Cfg.RealmdDB.Password,
    config.Cfg.RealmdDB.Hostname,
    config.Cfg.RealmdDB.Port,
    config.Cfg.RealmdDB.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realmd database (%v)\n", err)
    ApiDB.Close()
    return err
  }

  var db RealmDB
  for _, v := range config.Cfg.Cmangos.Realms {
    db.Name = v.Name
    db.Character, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
      v.Username, v.Password,
      v.Hostname, v.Port,
      v.Character))
    if err != nil {
      db.Character.Close()
      return err
    }
    db.World, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
      v.Username, v.Password,
      v.Hostname, v.Port,
      v.World))
    if err != nil {
      db.World.Close()
      return err
    }
    RealmsDB = append(RealmsDB, db)
  }

  return nil
}
