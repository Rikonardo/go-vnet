package vnet

import "net"

type builder struct {
	closeFunc func() error
	addrFunc  func() net.Addr
}

func Builder() *builder {
	return &builder{
		closeFunc: func() error { return nil },
		addrFunc:  func() net.Addr { return nil },
	}
}

func (b *builder) CloseFunc(f func() error) *builder {
	b.closeFunc = f
	return b
}

func (b *builder) AddrFunc(f func() net.Addr) *builder {
	b.addrFunc = f
	return b
}

func (b *builder) Build() *VirtualListener {
	return &VirtualListener{
		closeFunc: b.closeFunc,
		addrFunc:  b.addrFunc,
	}
}
