package realm

import (
  "metagit.org/blizzlike/cmangos-api/cmangos"
)

type Realm struct {
  Id int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Icon int `json:"icon,omitempty"`
  Population float64 `json:"population,omitempty"`
  Host cmangos.DaemonInfo `json:"host,omitempty"`
}
