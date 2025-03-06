package discovery

import (
	"fmt"
	"net"
	"testing"
)

func TestNetAddr(t *testing.T) {
	var n net.Addr = NewNetAddr("tcp", ":8090")
	//test.Assert(t, n.Network() == "tcp")
	//test.Assert(t, n.String() == "12345")

	fmt.Println(n.Network() == "tcp")
	fmt.Println(n.String() == ":8090")

	fmt.Println(Resolve("8090"))
	fmt.Println(Resolve(":8090"))
	fmt.Println(Resolve("localhost:8090"))
	fmt.Println(Resolve(n.String()))

	fmt.Println(LocalAddr()) // local ip
}
