package livestatus

import (
	"fmt"
	"testing"
	"time"
)

func Test_Command(t *testing.T) {
	expected := fmt.Sprintf("COMMAND [%d] command1\n\n", time.Now().Unix())

	c := NewCommand("command1")

	result := c.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_CommandWithInlineArgs(t *testing.T) {
	expected := fmt.Sprintf("COMMAND [%d] command1;arg1;arg2\n\n", time.Now().Unix())

	c := NewCommand("command1", "arg1", "arg2")

	result := c.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_CommandWithArgs(t *testing.T) {
	expected := fmt.Sprintf("COMMAND [%d] command1;arg1;arg2\n\n", time.Now().Unix())

	c := NewCommand("command1")
	c.Arg("arg1")
	c.Arg("arg2")

	result := c.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}
