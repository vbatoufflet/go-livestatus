package livestatus

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Command is a binding command instance.
type Command struct {
	cmd  string
	vals []string
	ls   *Livestatus
}

func newCommand(ls *Livestatus) *Command {
	return &Command{
		ls: ls,
	}
}

func (c *Command) Raw(cmd string) {
	c.cmd = cmd
}

func (c *Command) Arg(v interface{}) {
	c.vals = append(c.vals, fmt.Sprintf("%v", v))
}

type CommandOpFunc func(*Command)

func (c *Command) Op(op CommandOpFunc) {
	op(c)
}

// KeepAliveOff disables the default keepalive from Command
func (q *Query) KeepAliveOff() *Query {
	q.ls.keepalive = false
	return q
}

// Exec executes the query.
func (c *Command) Exec() (*Response, error) {
	resp := &Response{}

	var err error
	var conn net.Conn

	if c.ls.keepConn != nil {
		conn = c.ls.keepConn
	} else {
		// Connect to socket
		conn, err = c.dial()
		if err != nil {
			return nil, err
		}
	}

	if !c.ls.keepalive {
		c.ls.keepConn = nil
		defer conn.Close()
	} else {
		c.ls.keepConn = conn
	}

	// Send command data
	cmd, err := c.buildCmd(time.Now())
	if err != nil {
		return nil, err
	}

	conn.Write([]byte(cmd))
	// You get nothing back from an external command
	// no way of knowing if this has worked

	return resp, nil
}

func (c *Command) buildCmd(t time.Time) (string, error) {
	cmdStr := fmt.Sprintf("COMMAND [%d] %s", t.Unix(), c.cmd)
	cmdStr = fmt.Sprintf("%s;%s", cmdStr, strings.Join(c.vals, ";"))

	return fmt.Sprintf("%s\n", cmdStr), nil
}

func (c *Command) dial() (net.Conn, error) {
	if c.ls.dialer != nil {
		return c.ls.dialer()
	} else {
		return net.Dial(c.ls.network, c.ls.address)
	}
}
