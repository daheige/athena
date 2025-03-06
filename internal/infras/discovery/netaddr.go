package discovery

import (
	"fmt"
	"net"
	"strconv"
)

var _ net.Addr = &NetAddr{}

// NetAddr implements the net.Addr interface.
type NetAddr struct {
	network string
	address string
}

// NewNetAddr creates a new NetAddr object with the network and address provided.
func NewNetAddr(network, address string) net.Addr {
	return &NetAddr{network, address}
}

// Network implements the net.Addr interface.
func (na *NetAddr) Network() string {
	return na.network
}

// String implements the net.Addr interface.
func (na *NetAddr) String() string {
	return na.address
}

// Resolve returns host:port address
func Resolve(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", err
	}

	// if host is empty or "::", use local ipv4 address as host
	if host == "" || host == "::" {
		host, err = localIPv4Host()
		if err != nil {
			return "", fmt.Errorf("parse registry info addr error: %w", err)
		}
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return "", fmt.Errorf("parse address port error: %w", err)
	}

	return fmt.Sprintf("%s:%d", host, p), nil
}

// LocalAddr returns local ipv4 ip
func LocalAddr() (string, error) {
	return localIPv4Host()
}

func localIPv4Host() (string, error) {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range s {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}

	return "", fmt.Errorf("not found ipv4 address")
}
