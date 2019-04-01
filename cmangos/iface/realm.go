package iface

type Realmd struct {
  Address string `json:"address,omitempty"`
  Port int `json:"port,omitempty"`
  State int `json:"state"`
  Lastcheck int `json:"lastcheck"`
}

type Realm struct {
  Id int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Address string `json:"address,omitempty"`
  Port int `json:"port,omitempty"`
  Icon int `json:"icon,omitempty"`
  Population float64 `json:"population,omitempty"`
  State int `json:"state"`
  Lastcheck int `json:"lastcheck"`
}

type Realmlist struct {
  Realmd Realmd `json:"realmd,omitempty"`
  Realmlist []Realm `json:"realmlist"`
}
