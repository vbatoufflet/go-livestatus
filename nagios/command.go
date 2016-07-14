package nagios

import (
	"fmt"
	"time"

	lvst "github.com/vbatoufflet/go-livestatus"
)

//go:generate go run cmd/gen-nag-external-commands/main.go -o commands_generated.go

type stickyBool bool

func (b stickyBool) String() string {
	if b {
		return "2"
	} else {
		return "0"
	}
}

type normalBool bool

func (b normalBool) String() string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

type stringDuration struct{ time.Duration }

func (d stringDuration) String() string {
	return fmt.Sprintf("%d", d.Duration/time.Second)
}

func Hostname(c *lvst.Command, s string) {
	c.Arg(s)
}

func HostGroupName(c *lvst.Command, s string) {
	c.Arg(s)
}

func ServiceGroupName(c *lvst.Command, s string) {
	c.Arg(s)
}

func ContactName(c *lvst.Command, s string) {
	c.Arg(s)
}

func ContactGroupName(c *lvst.Command, s string) {
	c.Arg(s)
}

func ServiceDescription(c *lvst.Command, s string) {
	c.Arg(s)
}

func Sticky(c *lvst.Command, b bool) {
	c.Arg(stickyBool(b).String())
}

func Notify(c *lvst.Command, b bool) {
	c.Arg(normalBool(b).String())
}

func Fixed(c *lvst.Command, b bool) {
	c.Arg(normalBool(b).String())
}

func Persistent(c *lvst.Command, b bool) {
	c.Arg(normalBool(b).String())
}

func Delete(c *lvst.Command, b bool) {
	c.Arg(normalBool(b).String())
}

func Author(c *lvst.Command, s string) {
	c.Arg(s)
}

func Comment(c *lvst.Command, s string) {
	c.Arg(s)
}

func Start(c *lvst.Command, t time.Time) {
	c.Arg(t.Unix())
}

func End(c *lvst.Command, t time.Time) {
	c.Arg(t.Unix())
}

func CheckTime(c *lvst.Command, t time.Time) {
	c.Arg(t.Unix())
}

func NotificationTime(c *lvst.Command, t time.Time) {
	c.Arg(t.Unix())
}

func NotificationTimePeriod(c *lvst.Command, s string) {
	c.Arg(s)
}

func Duration(c *lvst.Command, t time.Duration) {
	c.Arg(stringDuration{t}.String())
}

func TriggerID(c *lvst.Command, i int) {
	c.Arg(i)
}

func DowntimeID(c *lvst.Command, i int) {
	c.Arg(i)
}

func CommentID(c *lvst.Command, i int) {
	c.Arg(i)
}

func Options(c *lvst.Command, i int) {
	c.Arg(i)
}

func CheckAttempts(c *lvst.Command, i int) {
	c.Arg(i)
}

func StatusCode(c *lvst.Command, i int) {
	c.Arg(i)
}

func ReturnCode(c *lvst.Command, i int) {
	c.Arg(i)
}

func NotificationNumber(c *lvst.Command, i int) {
	c.Arg(i)
}

func Value(c *lvst.Command, s string) {
	c.Arg(s)
}

func VarName(c *lvst.Command, s string) {
	c.Arg(s)
}

func VarValue(c *lvst.Command, s string) {
	c.Arg(s)
}

func EventHandlerCommand(c *lvst.Command, s string) {
	c.Arg(s)
}

func CheckCommand(c *lvst.Command, s string) {
	c.Arg(s)
}

func TimePeriod(c *lvst.Command, s string) {
	c.Arg(s)
}

func CheckTimePeriod(c *lvst.Command, s string) {
	c.Arg(s)
}

func FileName(c *lvst.Command, s string) {
	c.Arg(s)
}

func PluginOutput(c *lvst.Command, s string) {
	c.Arg(s)
}
