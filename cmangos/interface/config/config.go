package config

type InterfaceConfig struct {
  NeedInvite bool `json:"invite,omitempty"`
  RealmdAddress string `json:"realmd,omitempty"`
}
