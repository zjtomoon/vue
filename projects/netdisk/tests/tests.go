package tests

type Address struct {
	IP   string
	Port string
}

var address *Address

func SetAddress(ip string, port string) *Address {
	addr := &Address{}
	addr.IP = ip
	addr.Port = port
	address = addr
	return address
}
