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

// Query represents a Livestatus query instance.
type Query struct {
	table     string
	headers   []string
	columns   []string
	keepalive bool

	writeTimeout time.Duration
	readTimeout  time.Duration
}

// NewQuery creates a new Livestatus query instance.
func NewQuery(table string) *Query {
	return &Query{
		table:   table,
		headers: []string{},
		columns: []string{},
	}
}

// Columns selects which columns to retrieve.
func (q *Query) Columns(names ...string) *Query {
	q.headers = append(q.headers, "Columns: "+strings.Join(names, " "))
	q.columns = names
	return q
}

// Filter appends a new filter to the query.
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

// Limit sets the limit of datasets to retrieve.
func (q *Query) Limit(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("Limit: %d", n))
	return q
}

// WaitObject specifies an object from the table to wait on.
//
// For `hosts`, `hostgroups`, `servicegroups`, `contacts` and `contactgroups` tables this is simply the name of
// the object. For the `services` table it is the `hostname` and the service `description` separated by a space.
func (q *Query) WaitObject(name string) *Query {
	q.headers = append(q.headers, "WaitObject: "+name)
	return q
}

// WaitCondition appends a new wait condition to the query.
func (q *Query) WaitCondition(rule string) *Query {
	q.headers = append(q.headers, "WaitCondition: "+rule)
	return q
}

// WaitConditionAnd combines the n last wait conditions into a new wait condition using a `And` operation.
func (q *Query) WaitConditionAnd(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionAnd: %d", n))
	return q
}

// WaitConditionOr combines the n last wait conditions into a new wait condition using a `Or` operation.
func (q *Query) WaitConditionOr(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionOr: %d", n))
	return q
}

// WaitConditionNegate negates the most recent wait condition.
func (q *Query) WaitConditionNegate() *Query {
	q.headers = append(q.headers, "WaitConditionNegate:")
	return q
}

// WaitTrigger appends a new wait trigger to the query, waiting for a specific event broker message to recheck
// condition.
func (q *Query) WaitTrigger(event string) *Query {
	q.headers = append(q.headers, "WaitTrigger: "+event)
	return q
}

// WaitTimeout sets the upper limit on the time to wait before executing the query.
func (q *Query) WaitTimeout(d time.Duration) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitTimeout: %d", d/time.Millisecond))
	return q
}

// KeepAlive keeps the connection open to reuse for additional requests.
func (q *Query) KeepAlive() *Query {
	q.headers = append(q.headers, "KeepAlive: on")
	q.keepalive = true
	return q
}

// WriteTimeout sets the connection timeout for write operations.
// A value of 0 disables the timeout.
func (q *Query) WriteTimeout(timeout time.Duration) *Query {
	q.writeTimeout = timeout
	return q
}

// ReadTimeout sets the connection timeout for read operations.
// Be careful when using a read timeout in conjunction with wait conditions.
// A value of 0 disables the timeout.
func (q *Query) ReadTimeout(timeout time.Duration) *Query {
	q.readTimeout = timeout
	return q
}

// String returns a string representation of the Livestatus query.
func (q Query) String() string {
	s := "GET " + q.table
	if len(q.headers) > 0 {
		s += "\n" + strings.Join(q.headers, "\n")
	}
	s += "\nResponseHeader: fixed16\nOutputFormat: json\n\n"

	return s
}

func (q Query) handle(conn net.Conn) (*Response, error) {
	var err error

	cmd := q.String()
	lcmd := len(cmd)

	if q.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(q.writeTimeout))
	} else {
		// disable timeout
		conn.SetWriteDeadline(time.Time{})
	}

	// Send query data
	n, err := conn.Write([]byte(cmd))
	if err != nil {
		return nil, fmt.Errorf("sending query failed: %v", err)
	}

	if n != lcmd {
		return nil, fmt.Errorf("incomplete write to livestatus. Wrote %d bytes while %d were to be written", n, lcmd)
	}

	if q.readTimeout > 0 {
		conn.SetReadDeadline(time.Now().Add(q.readTimeout))
	} else {
		// disable timeout
		conn.SetReadDeadline(time.Time{})
	}

	// Read response header
	data := make([]byte, 16)

	_, err = conn.Read(data)
	if err != nil {
		return nil, fmt.Errorf("reading response header failed: %v", err)
	}

	resp := &Response{}
	resp.Status, err = strconv.Atoi(string(data[:3]))
	if err != nil {
		return nil, ParseError{
			Message:    fmt.Sprintf("parsing response status from header failed: %v", err),
			FailedData: data[:3],
			Buffer:     data,
		}
	}

	length, err := strconv.Atoi(string(bytes.TrimSpace(data[5:15])))
	if err != nil {
		return nil, ParseError{
			Message:    fmt.Sprintf("parsing response length from header failed: %v", err),
			FailedData: bytes.TrimSpace(data[5:15]),
			Buffer:     data,
		}
	}
	remainder := length

	// Receive response data
	buf := bytes.NewBuffer(nil)

	for {
		data = make([]byte, remainder)

		n, err := conn.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading body (buffer size: %d, remainder: %d) failed: %v", buf.Len(), remainder, err)
		}

		buf.Write(bytes.TrimRight(data, "\x00"))

		remainder -= n
		if remainder <= 0 {
			break
		}
	}

	if buf.Len() == 0 {
		return resp, nil
	}

	// Stop on invalid status
	if resp.Status >= 400 {
		resp.Message = strings.TrimRight(buf.String(), "\n")
		return resp, ErrInvalidQuery
	}

	// Parse received data for records
	resp.Records, err = q.parse(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("parsing read data as records failed: %v", err)
	}

	return resp, nil
}

func (q Query) keepAlive() bool {
	return q.keepalive
}

func (q *Query) parse(data []byte) ([]Record, error) {
	var rows [][]interface{}

	// Unmarshal received data
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, ParseError{
			Message:    fmt.Sprintf("unmarshalling JSON failed: %v", err),
			FailedData: data,
		}
	} else if len(q.columns) == 0 && len(rows) < 2 || len(q.columns) > 0 && len(rows) < 1 {
		return nil, nil
	}

	// Extract q.columns names from first row if no column provided
	if len(q.columns) == 0 {
		q.columns = make([]string, len(rows[0]))
		for i, value := range rows[0] {
			q.columns[i] = value.(string)
		}
		rows = rows[1:]
	}

	// Fill records maps
	records := []Record{}
	for _, row := range rows {
		r := Record{}
		for i, value := range row {
			r[q.columns[i]] = value
		}
		records = append(records, r)
	}

	return records, nil
}
