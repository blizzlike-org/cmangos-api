package realmd

import (
  "fmt"
  "net"
  "os"
  "time"

  "metagit.org/blizzlike/cmangos-api/cmangos/iface"

  "metagit.org/blizzlike/cmangos-api/modules/config"
)

func Check(r *iface.Realmd) error {
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

func GetRealmd() iface.Realmd {
  var r iface.Realmd
  r.Address = config.Cfg.Cmangos.Realmd
  r.Port = config.Cfg.Cmangos.Port
  Check(&r)
  return r
}
