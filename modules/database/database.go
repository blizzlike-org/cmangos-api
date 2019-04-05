package database

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

type MangosdDB struct {
  Id int
  Character *sql.DB
  World *sql.DB
}

var Api *sql.DB
var Realmd *sql.DB
var Mangosd []MangosdDB

func Close() {
  Api.Close()
  Realmd.Close()

  for _, v := range Mangosd {
    v.Character.Close()
    v.World.Close()
  }
}

func Open() error {
  var err error
  Api, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    config.Settings.Api.Db.Username,
    config.Settings.Api.Db.Password,
    config.Settings.Api.Db.Address,
    config.Settings.Api.Db.Port,
    config.Settings.Api.Db.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to api database (%v)\n", err)
    return err
  }

  Realmd, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
    config.Settings.Realmd.Db.Username,
    config.Settings.Realmd.Db.Password,
    config.Settings.Realmd.Db.Address,
    config.Settings.Realmd.Db.Port,
    config.Settings.Realmd.Db.Database))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realmd database (%v)\n", err)
    Api.Close()
    return err
  }

  var db MangosdDB
  for _, v := range config.Settings.Mangosd {
    db.Id = v.Id
    db.Character, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
      v.CharacterDb.Username, v.CharacterDb.Password,
      v.CharacterDb.Address, v.CharacterDb.Port,
      v.CharacterDb.Database))
    if err != nil {
      db.Character.Close()
      return err
    }
    db.World, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
      v.WorldDb.Username, v.WorldDb.Password,
      v.WorldDb.Address, v.WorldDb.Port,
      v.WorldDb.Database))
    if err != nil {
      db.World.Close()
      return err
    }
    Mangosd = append(Mangosd, db)
  }

  return nil
}
