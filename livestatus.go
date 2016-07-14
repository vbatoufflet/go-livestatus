// Package livestatus provides a binding to MK Livestatus sockets.
package livestatus

import "net"

// Livestatus is a binding instance.
type Livestatus struct {
	network string
	address string
	dialer  func() (net.Conn, error)

	keepalive bool
	keepConn  net.Conn
}

// Close any open connection from a KeepAlive
func (l *Livestatus) Close() error {
	l.keepalive = false
	if l.keepConn != nil {
		l.keepConn = nil
		return l.keepConn.Close()
	}
	return nil
}

// Query creates a new query instance on a spacific table.
func (l *Livestatus) Query(table string) *Query {
	l.keepalive = false
	return newQuery(table, l)
}

// Command creates a new command instanc.
func (l *Livestatus) Command() *Command {
	l.keepalive = true
	return newCommand(l)
}

// NewLivestatus creates a new binding instance.
func NewLivestatus(network, address string) *Livestatus {
	return &Livestatus{
		network: network,
		address: address,
	}
}

// NewLivestatusWithDialer creates a new binding that uses the net.Conn returned
// by the provided dialer function.
func NewLivestatusWithDialer(dialer func() (net.Conn, error)) *Livestatus {
	return &Livestatus{
		dialer: dialer,
	}
}
