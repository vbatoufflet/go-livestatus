package livestatus

import (
	"reflect"
	"testing"
)

func Test_Query(t *testing.T) {
	expected := "GET table1\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryColumnSingle(t *testing.T) {
	expected := "GET table1\n"
	expected += "Columns: column1\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Columns("column1")

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryColumnMulti(t *testing.T) {
	expected := "GET table1\n"
	expected += "Columns: column1 column2\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Columns("column1", "column2")

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryFilterSingle(t *testing.T) {
	expected := "GET table1\n"
	expected += "Filter: column1 ~ abc\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Filter("column1 ~ abc")

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryFilterMulti(t *testing.T) {
	expected := "GET table1\n"
	expected += "Filter: column1 ~ abc\n"
	expected += "Filter: column2 >= 123\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryAnd(t *testing.T) {
	expected := "GET table1\n"
	expected += "Filter: column1 ~ abc\n"
	expected += "Filter: column2 >= 123\n"
	expected += "And: 2\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")
	q.And(2)

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryOr(t *testing.T) {
	expected := "GET table1\n"
	expected += "Filter: column1 ~ abc\n"
	expected += "Filter: column2 >= 123\n"
	expected += "Or: 2\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Filter("column1 ~ abc")
	q.Filter("column2 >= 123")
	q.Or(2)

	result := q.buildCmd()
	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}

func Test_QueryNegate(t *testing.T) {
	expected := "GET table1\n"
	expected += "Filter: column1 ~ abc\n"
	expected += "Negate:\n"
	expected += "ResponseHeader: fixed16\n"
	expected += "OutputFormat: json\n"

	q := newQuery("table1", &Livestatus{})
	q.Filter("column1 ~ abc")
	q.Negate()

	result := q.buildCmd()
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
		Record{"name": "name1", "value": 123.0},
		Record{"name": "name2", "value": 456.0},
	}

	q := newQuery("table1", &Livestatus{})

	result, err := q.parse([]byte(data))
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}

func Test_QueryParseColumns(t *testing.T) {
	data := `[
		["name1", 123],
		["name2", 456]
	]`

	expected := []Record{
		Record{"name": "name1", "value": 123.0},
		Record{"name": "name2", "value": 456.0},
	}

	q := newQuery("table1", &Livestatus{})
	q.Columns("name", "value")

	result, err := q.parse([]byte(data))
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}
