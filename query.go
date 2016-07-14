package livestatus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

// Query is a binding query instance.
type Query struct {
	table   string
	headers []string
	columns []string
	ls      *Livestatus
}

// Columns sets the names of the columns to retrieve when executing a query.
func (q *Query) Columns(names ...string) *Query {
	q.headers = append(q.headers, "Columns: "+strings.Join(names, " "))
	q.columns = names
	return q
}

// KeepAlive keeps the connection open after the query, for re-use
func (q *Query) KeepAlive() *Query {
	q.headers = append(q.headers, "KeepAlive: on")
	return q
}

// Filter sets a new filter to apply to the query.
func (q *Query) Filter(rule string) *Query {
	q.headers = append(q.headers, "Filter: "+rule)
	return q
}

// And combines the n last filters into a new filter using a `And` operation.
func (q *Query) And(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("And: %d", n))
	return q
}

// Or combines the n last filters into a new filter using a `Or` operation.
func (q *Query) Or(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("Or: %d", n))
	return q
}

// Negate negates the most recent filter.
func (q *Query) Negate() *Query {
	q.headers = append(q.headers, "Negate:")
	return q
}

// Limit the query to n responses.
func (q *Query) Limit(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("Limit: %d", n))
	return q
}

// WaitObject sets the object within the queried table to wait on. For the table
// hosts, hostgroups, servicegroups, contacts and contactgroups this is simply
// the name of the object. For the table services it is the hostname followed
// by a space followed by the service description
func (q *Query) WaitObject(obj string) *Query {
	q.headers = append(q.headers, "WaitObject: "+obj)
	return q
}

// WaitCondition sets a new wait condition  to apply to the query.
func (q *Query) WaitCondition(rule string) *Query {
	q.headers = append(q.headers, "WaitCondition: "+rule)
	return q
}

// WaitConditionAnd combines the n last wait conditions into a new wait
// condition using a `And` operation.
func (q *Query) WaitConditionAnd(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionAnd: %d", n))
	return q
}

// WaitConditionOr combines the n last wait condition into a new wait condition
// using a `Or` operation.
func (q *Query) WaitConditionOr(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionOr: %d", n))
	return q
}

// WaitConditionNegate negates the most recent wait condition.
func (q *Query) WaitConditionNegate() *Query {
	q.headers = append(q.headers, "WaitConditionNegate:")
	return q
}

// WaitTrigger sets the nagios event that will trigger a check of
// the wait condition.
func (q *Query) WaitTrigger(event string) *Query {
	q.headers = append(q.headers, "WaitTrigger: "+event)
	return q
}

// WaitTimeout set a timeout for the wait condition.
func (q *Query) WaitTimeout(t time.Duration) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitTimeout: %d", t/time.Millisecond))
	return q
}

// Exec executes the query.
func (q *Query) Exec() (*Response, error) {
	resp := &Response{}

	var err error
	var conn net.Conn

	if q.ls.keepConn != nil {
		conn = q.ls.keepConn
	} else {
		// Connect to socket
		conn, err = q.dial()
		if err != nil {
			return nil, err
		}
	}

	if !q.ls.keepalive {
		q.ls.keepConn = nil
		defer conn.Close()
	} else {
		q.ls.keepConn = conn
	}

	// Send command data
	conn.Write([]byte(q.buildCmd()))

	// Read response header
	data := make([]byte, 16)

	_, err = conn.Read(data)
	if err != nil {
		return nil, err
	}

	resp.Status, err = strconv.Atoi(string(data[:3]))
	if err != nil {
		return nil, err
	}

	// Receive response data
	buf := bytes.NewBuffer(nil)

	for {
		data = make([]byte, 1024)

		n, err := conn.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		buf.Write(bytes.TrimRight(data, "\x00"))

		// New line signals the end of content. This check helps
		// if the connection is not forcibly closed
		if data[n-1] == byte('\n') {
			break
		}
	}

	if buf.Len() == 0 {
		return resp, nil
	}

	// Parse received data for records
	resp.Records, err = q.parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *Query) buildCmd() string {
	cmd := "GET " + q.table

	// Append headers if any
	if len(q.headers) > 0 {
		cmd += "\n" + strings.Join(q.headers, "\n")
	}

	// Set default headers
	cmd += "\nResponseHeader: fixed16"
	cmd += "\nOutputFormat: json"
	cmd += "\n\n"

	return cmd
}

func (q *Query) dial() (net.Conn, error) {
	if q.ls.dialer != nil {
		return q.ls.dialer()
	} else {
		return net.Dial(q.ls.network, q.ls.address)
	}
}

func (q *Query) parse(data []byte) ([]Record, error) {
	var (
		records []Record
		rows    [][]interface{}
	)

	// Unmarshal received data
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	} else if len(q.columns) == 0 && len(rows) < 2 || len(q.columns) > 0 && len(rows) < 1 {
		return records, nil
	}

	// Extract columns names from first row
	start := 0

	if len(q.columns) == 0 {
		q.columns = make([]string, len(rows[0]))
		for i, value := range rows[0] {
			q.columns[i] = value.(string)
		}

		start = 1
	}

	// Fill records maps
	for _, row := range rows[start:] {
		r := make(Record)
		for i, value := range row {
			r.set(q.columns[i], value)
		}
		records = append(records, r)
	}

	return records, nil
}

func newQuery(table string, ls *Livestatus) *Query {
	return &Query{
		table:   table,
		headers: make([]string, 0),
		ls:      ls,
	}
}
