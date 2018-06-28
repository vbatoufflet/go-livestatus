package livestatus

import (
	"net"
)

// Client represents a Livestatus client instance.
type Client struct {
	network string
	address string
	dialer  *net.Dialer
	conn    net.Conn
}

// NewClient creates a new Livestatus client instance.
func NewClient(network, address string) *Client {
	return NewClientWithDialer(network, address, new(net.Dialer))
}

// NewClientWithDialer creates a new Livestatus client instance using a provided network dialer.
func NewClientWithDialer(network, address string, dialer *net.Dialer) *Client {
	return &Client{
		network: network,
		address: address,
		dialer:  dialer,
	}
}

// Close closes any remaining connection.
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}

// Exec executes a given Livestatus query.
func (c *Client) Exec(r Request) (*Response, error) {
	var err error

	// Initialize connection if none available
	if c.conn == nil {
		c.conn, err = c.dialer.Dial(c.network, c.address)
		if err != nil {
			return nil, err
		}

		if r.keepAlive() {
			switch c.network {
			case "tcp":
				c.conn.(*net.TCPConn).SetKeepAlive(true)
			}
		} else {
			defer c.Close()
		}
	}

	return r.handle(c.conn)
}
