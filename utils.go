package fastsocket

import (
	"net"
	"syscall"
)

func IsNetTemporary(err error) bool {
	if err == syscall.EAGAIN {
		return true
	}
	if ne, ok := err.(net.Error); ok && ne.Temporary() {
		return true
	}
	return false
}

func resolveSockAddr4(netaddr string) (syscall.Sockaddr, error) {
	addr, err := net.ResolveTCPAddr("tcp4", netaddr)
	if err != nil {
		return nil, err
	}
	ip := addr.IP
	if len(ip) == 0 {
		ip = net.IPv4zero
	}
	sa4 := &syscall.SockaddrInet4{Port: addr.Port}
	copy(sa4.Addr[:], ip.To4())
	return sa4, nil
}