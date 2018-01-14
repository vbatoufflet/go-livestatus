package nagios

import (
	"fmt"
	"time"
)

//go:generate go run internal/generate-commands/main.go -o commands.go

func stringifyArg(name, typ string, v interface{}) string {
	switch typ {
	case "bool":
		if !v.(bool) {
			return "0"
		} else if name == "sticky" {
			return "2"
		}

		return "1"

	case "int", "interface{}", "string":
		return fmt.Sprintf("%v", v)

	case "time.Duration":
		return fmt.Sprintf("%d", v.(time.Duration)/time.Second)

	case "time.Time":
		return fmt.Sprintf("%d", v.(time.Time).Unix())
	}

	return ""
}
