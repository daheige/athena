package discovery

import (
	"errors"
	"net"
	"strconv"
	"strings"
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

// Resolve 解析address地址
func Resolve(network, address string) (string, error) {
	if network == "tcp" {
		s := strings.Split(address, ":")
		switch len(s) {
		case 0:
			return "", errors.New("address is empty")
		case 1:
			_, err := strconv.Atoi(s[0])
			if err != nil {
				return "", errors.New("invalid port")
			}
			address = net.JoinHostPort("0.0.0.0", s[0])
		case 2:
			if s[0] == "" {
				s[0] = "0.0.0.0"
			}

			_, err := strconv.Atoi(s[1])
			if err != nil {
				return "", errors.New("invalid port")
			}
			address = net.JoinHostPort(s[0], s[1])
		default:
			return "", errors.New("invalid address")
		}
	}

	// check address
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return "", err
	}

	return net.JoinHostPort(addr.IP.String(), strconv.Itoa(addr.Port)), nil
}

// LocalAddr 获取本地ip地址
func LocalAddr() (string, error) {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range s {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}

	return "", errors.New("no local address")
}
