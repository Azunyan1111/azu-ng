package GetGateway

import (
	"github.com/jackpal/gateway"
	"net"
)

func GetDefaultGatewayForInterface()string{
	ip,err := gateway.DiscoverGateway()
	if err != nil{
		panic(err)
	}
	return ip.String()
}

func GetDefaultGatewayForIPAddress(str string)(addr string){
	// Check ip addr
	ip := net.ParseIP(str)
	if ip.String() == "<nil>"{
		return ""
	}
	return ip.String()
}
