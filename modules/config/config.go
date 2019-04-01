package config

import (
  ini "gopkg.in/ini.v1"
)

type ConfigDB struct {
  Username string
  Password string
  Hostname string
  Port int
  Database string
}

type ConfigCmangos struct {
  Realmd string
  Port int
  Timeout int
  Interval int
}

type Config struct {
  Listen string
  Port int
  ApiDB ConfigDB
  RealmdDB ConfigDB
  NeedInvite bool
  Cmangos ConfigCmangos
}

var Cfg Config

func Read(file string) (Config, error) {
  c, err := ini.Load(file)
  if err != nil {
    return Cfg, err
  }

  Cfg.Listen = c.Section("server").Key("listen").MustString("127.0.0.1")
  Cfg.Port = c.Section("server").Key("port").MustInt(5556)

  Cfg.ApiDB.Hostname = c.Section("apidb").Key("hostname").MustString("127.0.0.1")
  Cfg.ApiDB.Port = c.Section("apidb").Key("port").MustInt(3306)
  Cfg.ApiDB.Username = c.Section("apidb").Key("username").MustString("cmangos-api")
  Cfg.ApiDB.Password = c.Section("apidb").Key("password").MustString("cmangos-api")
  Cfg.ApiDB.Database = c.Section("apidb").Key("database").MustString("cmangos-api")

  Cfg.RealmdDB.Hostname = c.Section("realmddb").Key("hostname").MustString("127.0.0.1")
  Cfg.RealmdDB.Port = c.Section("realmddb").Key("port").MustInt(3306)
  Cfg.RealmdDB.Username = c.Section("realmddb").Key("username").MustString("mangos")
  Cfg.RealmdDB.Password = c.Section("realmddb").Key("password").MustString("mangos")
  Cfg.RealmdDB.Database = c.Section("realmddb").Key("database").MustString("realmd")

  Cfg.NeedInvite = c.Section("account").Key("needInvite").MustBool(false)

  Cfg.Cmangos.Realmd = c.Section("cmangos").Key("realmd").MustString("logon.example.org")
  Cfg.Cmangos.Port = c.Section("cmangos").Key("port").MustInt(3724)
  Cfg.Cmangos.Timeout = c.Section("cmangos").Key("timeout").MustInt(10)
  Cfg.Cmangos.Interval = c.Section("cmangos").Key("interval").MustInt(300)

  return Cfg, nil
}
