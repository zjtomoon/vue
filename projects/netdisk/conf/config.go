package conf

import (
	"encoding/json"
	"os"
)

type Config struct {
	WebAddr         string `json:"WebAddr"`
	WebIndex        string `json:"WebIndex"`
	StaticFS        string `json:"StaticFS"`
	FilePath        string `json:"FilePath"`
	FileDiskTotal   uint64 `json:"FileDiskTotal"`
	SaveFileMultipe bool   `json:"SaveFileMultipe"`
	Username        string `json:"Username"`
	Password        string `json:"Password"`
}

var config *Config

func LoadConfig(path string) *Config {
	file,_ := os.Open(path)
	decoder := json.NewDecoder(file)
	conf := &Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}