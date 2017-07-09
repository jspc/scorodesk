package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Listen string `json:"listen_address"`
	Scoro  ScoroConfig
}

type ScoroConfig struct {
	Company string
	Key     string
}

func LoadConfig(f string) (c Config, err error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &c)
	return
}
