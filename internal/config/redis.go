package config

import "time"

type redis struct {
	Address      string        `json:"address,omitempty" default:"address"`
	Password     string        `json:"password,omitempty" default:"password"`
	DB           int           `json:"db,omitempty" default:"db"`
	DialTimeout  time.Duration `json:"dial_timeout,omitempty" default:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout,omitempty" default:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty" default:"write_timeout"`
}

var Redis *redis
