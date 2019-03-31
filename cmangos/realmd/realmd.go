package realmd

import (
  "fmt"
  "net"
  "os"
  "time"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

type Realmd struct {
  Address string `json:"address,omitempty"`
  Port int `json:"port,omitempty"`
  State int `json:"state"`
  Lastcheck int `json:"lastcheck"`
}

func (r *Realmd) Check() error {
  c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", r.Address, r.Port))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to realm %s:%d (%v)\n", r.Address, r.Port, err)
    r.State = 0
    return err
  }
  defer c.Close()

  r.State = 1
  r.Lastcheck = int(time.Now().Unix())
  return nil
}

func GetRealmd() Realmd {
  var r Realmd
  r.Address = config.Cfg.Cmangos.Realmd
  r.Port = config.Cfg.Cmangos.Port
  r.Check()
  return r
}
