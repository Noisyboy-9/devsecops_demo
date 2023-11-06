package config

type redis struct {
	Address  string `json:"address,omitempty" default:"address"`
	Password string `json:"password,omitempty" default:"password"`
	DB       int    `json:"db,omitempty" default:"db"`
}

var Redis *redis
