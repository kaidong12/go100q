package net

import (
	"fmt"
	"net"
)

func Net_demo_To4() {
	ipv6 := net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	//ipv6 := net.ParseIP("192.168.1.1")
	ipv4 := net.IP.To4(ipv6)

	if ipv4 != nil {
		fmt.Println("IPv4 address:", ipv4)
	} else {
		fmt.Println("IPv6 address has no IPv4 representation.")
	}
}
