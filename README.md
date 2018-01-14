go-livestatus
=============

[![][travis-badge]][travis-url] [![][godoc-badge]][godoc-url] [![][report-badge]][report-url]

This package implements a [MK Livestatus][mklivestatus-url] binding for Go.

The source code is available at [Github][project-url], licensed under the terms of the [BSD license][license-url].

Usage
-----

```go
package main

import (
	"fmt"
	"os"

	livestatus "github.com/vbatoufflet/go-livestatus"
)

func main() {
	c := livestatus.NewClient("tcp", "localhost:6557")
	// or c := livestatus.NewClient("unix", "/var/run/nagios/livestatus.sock")
	defer c.Close()

	q := livestatus.NewQuery("hosts")
	q.Columns("name", "state", "last_time_down")
	q.Filter("name ~ ^db[0-9]+\\.")
	// or q := livestatus.Query("hosts").Columns("name", "state", "last_time_down").Filter("name ~ ^db[0-9]+\\.")

	resp, err := c.Exec(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}

	for _, r := range resp.Records {
		name, err := r.GetString("name")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		state, err := r.GetInt("state")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		lastTimeDown, err := r.GetTime("last_time_down")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		fmt.Printf("Host: %s, State: %d, Last time down: %s\n", name, state, lastTimeDown)
	}
}
```

Output example:

```
Host: db1.example.net, State: 0, Last time down: 2015-04-03 06:54:32 +0200 CEST
Host: db2.example.net, State: 0, Last time down: 2015-06-07 12:34:56 +0200 CEST
```

[godoc-badge]: https://godoc.org/github.com/vbatoufflet/go-livestatus?status.svg
[godoc-url]: https://godoc.org/github.com/vbatoufflet/go-livestatus
[license-url]: https://opensource.org/licenses/BSD-3-Clause
[mklivestatus-url]: https://mathias-kettner.com/checkmk_livestatus.html
[project-url]: https://github.com/vbatoufflet/go-livestatus
[report-badge]: https://goreportcard.com/badge/github.com/vbatoufflet/go-livestatus
[report-url]: https://goreportcard.com/report/github.com/vbatoufflet/go-livestatus
[travis-badge]: https://travis-ci.org/vbatoufflet/go-livestatus.svg
[travis-url]: https://travis-ci.org/vbatoufflet/go-livestatus
