package inject // github.com/guesslin/inject

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// TODO:
// ./app --scrIP <ip> --srcMac <mac> --srcPort <port> --dstIP <IP> --dstMac <mac> --dstPort <port>

type Endpoint struct {
	IP   net.IP
	Mac  net.HardwareAddr
	Port uint16
}

func NewEndpoint(host string, rawMac string) (*Endpoint, error) {
	mac, err := net.ParseMAC(rawMac)
	if err != nil {
		return nil, err
	}
	elements := strings.Split(host, ":")
	var port uint16 = 80
	var ip net.IP
	if len(elements) >= 2 {
		p, err := strconv.ParseUint(elements[1], 10, 64)
		if err != nil {
			return nil, err
		}
		port = uint16(p)
	}
	ip = net.ParseIP(elements[0])
	if ip == nil {
		return nil, fmt.Errorf("host is an invalid ip")
	}
	return &Endpoint{
		IP:   ip,
		Mac:  mac,
		Port: port,
	}, nil
}

type Packet struct {
	Source      Endpoint
	Destination Endpoint
}
