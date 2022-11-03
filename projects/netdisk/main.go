package main

import (
	"flag"
	"netdisk/service"
)

//var address *tests.Address

var (
	file = flag.String("file", "./config.json", "config file")
)

func main() {
	//address = tests.SetAddress("127.0.0.1","8080")
	//fmt.Println(address.IP)
	//fmt.Println(address.Port)
	flag.Parse()

	logger := service.InitLogger("log", "filecloud")
	service.LoadConfig(*file)

	service.Launch()

	logger.Infof("stop")

}
