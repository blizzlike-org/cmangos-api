package database

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

var ApiDB *sql.DB
var RealmdDB *sql.DB

func Close() {
  ApiDB.Close()
  RealmdDB.Close()
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

  return nil
}
