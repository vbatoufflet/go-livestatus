package livestatus

import "net"

// Request represents Livestatus request interface.
type Request interface {
	String() string

	handle(net.Conn) (*Response, error)
	keepAlive() bool
}
