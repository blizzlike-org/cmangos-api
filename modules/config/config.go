package config

import (
  ini "gopkg.in/ini.v1"
)

var cfg *ini.File

type ConfigDB struct {
  Username string
  Password string
  Hostname string
  Port int
  Database string
}

type ConfigCmangos struct {
  Realmd string
}

type Config struct {
  Listen string
  Port int
  ApiDB ConfigDB
  RealmdDB ConfigDB
  NeedInvite bool
  Cmangos ConfigCmangos
}

func Read(file string) (Config, error) {
  var cfg Config
  c, err := ini.Load(file)
  if err != nil {
    return cfg, err
  }

  cfg.Listen = c.Section("server").Key("listen").MustString("127.0.0.1")
  cfg.Port = c.Section("server").Key("port").MustInt(5556)

  cfg.ApiDB.Hostname = c.Section("apidb").Key("hostname").MustString("127.0.0.1")
  cfg.ApiDB.Port = c.Section("apidb").Key("port").MustInt(3306)
  cfg.ApiDB.Username = c.Section("apidb").Key("username").MustString("cmangos-api")
  cfg.ApiDB.Password = c.Section("apidb").Key("password").MustString("cmangos-api")
  cfg.ApiDB.Database = c.Section("apidb").Key("database").MustString("cmangos-api")

  cfg.RealmdDB.Hostname = c.Section("realmddb").Key("hostname").MustString("127.0.0.1")
  cfg.RealmdDB.Port = c.Section("realmddb").Key("port").MustInt(3306)
  cfg.RealmdDB.Username = c.Section("realmddb").Key("username").MustString("mangos")
  cfg.RealmdDB.Password = c.Section("realmddb").Key("password").MustString("mangos")
  cfg.RealmdDB.Database = c.Section("realmddb").Key("database").MustString("realmd")

  cfg.NeedInvite = c.Section("account").Key("needInvite").MustBool(false)

  cfg.Cmangos.Realmd = c.Section("cmangos").Key("realmd").MustString("logon.example.org")

  return cfg, nil
}
