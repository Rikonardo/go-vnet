package vnet

import (
	"net"
)

type VirtualListener struct {
	connections []net.Conn
	closeFunc   func() error
	addrFunc    func() net.Addr
}

func (l *VirtualListener) Accept() (net.Conn, error) {
	for {
		if len(l.connections) > 0 {
			con := l.connections[len(l.connections)-1]
			l.connections = removeConn(l.connections)
			return con, nil
		}
	}
}

func removeConn(s []net.Conn) []net.Conn {
	return s[:len(s)-1]
}

func (l *VirtualListener) Close() error {
	for _, c := range l.connections {
		c.Close()
	}
	return l.closeFunc()
}

func (l *VirtualListener) Addr() net.Addr {
	return l.addrFunc()
}

func (l *VirtualListener) ConnectExisting(c net.Conn) {
	l.connections = append(l.connections, c)
}

func (l *VirtualListener) Connect() net.Conn {
	c1, c2 := net.Pipe()
	l.ConnectExisting(c1)
	return c2
}
