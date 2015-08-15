go-livestatus [![](https://api.travis-ci.org/vbatoufflet/go-livestatus.png)](https://travis-ci.org/vbatoufflet/go-livestatus)
=============

This package implements a MK Livestatus binding for Go.

The source code is available at [Github][0], licensed under the terms of the [BSD license][1].

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
	l := livestatus.NewLivestatus("tcp", "192.168.0.10:6557")
	// or l := livestatus.NewLivestatus("unix", "/var/run/nagios/livestatus.sock")

	q := l.Query("hosts")
	q.Columns("name", "state", "last_time_down")
	q.Filter("name ~ ^db[0-9]+\\.")
	// or q := l.Query("hosts").Columns("name", "state", "last_time_down").Filter("name ~ ^db[0-9]+\\.")

	resp, err := q.Exec()
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

		lastDown, err := r.GetTime("last_time_down")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		fmt.Printf("Host: %s, State: %d, Last time down: %s\n", name, state, lastDown)
	}
}
```

Output example:

```
Host: db1.example.net, State: 0, Last time down: 2015-04-03 06:54:32 +0200 CEST
Host: db2.example.net, State: 0, Last time down: 2015-06-07 12:34:56 +0200 CEST
```

See [GoDoc](https://godoc.org/github.com/vbatoufflet/go-livestatus) for further details.

[0]: https://github.com/vbatoufflet/go-livestatus
[1]: http://opensource.org/licenses/BSD-3-Clause
