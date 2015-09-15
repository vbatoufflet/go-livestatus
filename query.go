package livestatus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
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

// Exec executes the query.
func (q *Query) Exec() (*Response, error) {
	resp := &Response{}

	// Connect to socket
	conn, err := q.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

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

		_, err = conn.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		buf.Write(bytes.TrimRight(data, "\x00"))
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
	return net.Dial(q.ls.network, q.ls.address)
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
