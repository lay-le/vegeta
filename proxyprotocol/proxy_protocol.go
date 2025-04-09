package proxyprotocol

import (
	"context"
	"net"
	"time"

	"github.com/pires/go-proxyproto"
)

type ProxyDialer struct {
	net.Dialer
	ProxyHeader *proxyproto.Header
}

func NewProxyDialer() ProxyDialer {
	return ProxyDialer{
		Dialer: net.Dialer{
			Timeout:   time.Second * 3,
			KeepAlive: time.Second * 30,
		},
	}
}

func (p *ProxyDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := p.Dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	proto := proxyproto.TCPv4
	if conn.RemoteAddr().Network() == "tcp6" {
		proto = proxyproto.TCPv6
	}
	p.ProxyHeader = &proxyproto.Header{
		Version:           1,
		Command:           proxyproto.PROXY,
		TransportProtocol: proto, // how to check if addr is v4 or v5?
		SourceAddr:        conn.LocalAddr(),
		DestinationAddr:   conn.RemoteAddr(),
	}
	// CRLF is appened to the last.
	_, err = p.ProxyHeader.WriteTo(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
