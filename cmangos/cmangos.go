package cmangos

import (
  "fmt"
  "net"
  "os"
  "time"
)

type DaemonInfo struct {
  Address string `json:"address,omitempty"`
  Port int `json:"port,omitempty"`
  State int `json:"state"`
  Lastcheck int `json:"lastcheck"`
}

func CheckDaemon(d *DaemonInfo, timeout time.Duration) error {
  d.Lastcheck = int(time.Now().Unix())
  dialer := net.Dialer{Timeout: timeout * time.Second}
  c, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", d.Address, d.Port))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Cannot connect to daemon %s:%d (%v)\n",
      d.Address, d.Port, err)
    d.State = 0
    return err
  }
  c.Close()

  d.State = 1
  return nil
}
