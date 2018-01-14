package livestatus

import (
	"reflect"
	"testing"
	"time"
)

func Test_Query(t *testing.T) {
	expected := `GET table1
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryColumnsSingle(t *testing.T) {
	expected := `GET table1
Columns: column1
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Columns("column1")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryColumnsMulti(t *testing.T) {
	expected := `GET table1
Columns: column1 column2
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Columns("column1", "column2")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryFilterSingle(t *testing.T) {
	expected := `GET table1
Filter: column1 ~ abc
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Filter("column1 ~ abc")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryFilterMulti(t *testing.T) {
	expected := `GET table1
Filter: column1 ~ abc
Filter: column2 >= 123
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryAnd(t *testing.T) {
	expected := `GET table1
Filter: column1 ~ abc
Filter: column2 >= 123
And: 2
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")
	q.And(2)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryOr(t *testing.T) {
	expected := `GET table1
Filter: column1 ~ abc
Filter: column2 >= 123
Or: 2
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")
	q.Or(2)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryNegate(t *testing.T) {
	expected := `GET table1
Filter: column1 ~ abc
Negate:
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Filter("column1 ~ abc")
	q.Negate()

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryLimit(t *testing.T) {
	expected := `GET table1
Limit: 3
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.Limit(3)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryKeepAlive(t *testing.T) {
	expected := `GET table1
KeepAlive: on
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.KeepAlive()

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitObject(t *testing.T) {
	expected := `GET table1
WaitObject: object1
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitObject("object1")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitConditionSingle(t *testing.T) {
	expected := `GET table1
WaitCondition: column1 ~ abc
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitCondition("column1 ~ abc")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitConditionMulti(t *testing.T) {
	expected := `GET table1
WaitCondition: column1 ~ abc
WaitCondition: column2 >= 123
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitCondition("column1 ~ abc")
	q.WaitCondition("column2 >= 123")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitConditionAnd(t *testing.T) {
	expected := `GET table1
WaitCondition: column1 ~ abc
WaitCondition: column2 >= 123
WaitConditionAnd: 2
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitCondition("column1 ~ abc")
	q.WaitCondition("column2 >= 123")
	q.WaitConditionAnd(2)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitConditionOr(t *testing.T) {
	expected := `GET table1
WaitCondition: column1 ~ abc
WaitCondition: column2 >= 123
WaitConditionOr: 2
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitCondition("column1 ~ abc")
	q.WaitCondition("column2 >= 123")
	q.WaitConditionOr(2)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitConditionNegate(t *testing.T) {
	expected := `GET table1
WaitCondition: column1 ~ abc
WaitConditionNegate:
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitCondition("column1 ~ abc")
	q.WaitConditionNegate()

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitTrigger(t *testing.T) {
	expected := `GET table1
WaitTrigger: event
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitTrigger("event")

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryWaitTimeout(t *testing.T) {
	expected := `GET table1
WaitTimeout: 5000
ResponseHeader: fixed16
OutputFormat: json

`

	q := NewQuery("table1")
	q.WaitTimeout(5 * time.Second)

	result := q.String()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryParse(t *testing.T) {
	data := `[
		["name", "value"],
		["name1", 123],
		["name2", 456]
	]`

	expected := []Record{
		{"name": "name1", "value": 123.0},
		{"name": "name2", "value": 456.0},
	}

	q := NewQuery("table1")

	result, err := q.parse([]byte(data))
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}

func Test_QueryParseWithColumns(t *testing.T) {
	data := `[
		["name1", 123],
		["name2", 456]
	]`

	expected := []Record{
		{"name": "name1", "value": 123.0},
		{"name": "name2", "value": 456.0},
	}

	q := NewQuery("table1")
	q.Columns("name", "value")

	result, err := q.parse([]byte(data))
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}
