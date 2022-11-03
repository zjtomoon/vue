package service

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := LoadConfig("../config.json")
	fmt.Println(conf.WebAddr)
	fmt.Println(conf)
}
