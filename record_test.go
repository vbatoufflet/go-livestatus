package livestatus

import (
	"reflect"
	"testing"
	"time"
)

func Test_RecordLen(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	if record.Len() != 2 {
		t.Logf("\nExpected 2\nbut got  %#v\n", record.Len())
		t.Fail()
	}
}

func Test_RecordColumns(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	expected := []string{"name", "value"}

	result := record.Columns()
	if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}

func Test_RecordGet(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	expected := interface{}("name1")

	result, err := record.Get("name")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}

	expected = interface{}(123.0)

	result, err = record.Get("value")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}

func Test_RecordGetBool(t *testing.T) {
	record := Record{
		"name":   "name1",
		"active": 0.0,
	}

	result, err := record.GetBool("active")
	if err != nil {
		t.Fatal(err)
	} else if result {
		t.Logf("\nExpected %#v\nbut got  %#v\n", false, result)
		t.Fail()
	}

	record = Record{
		"name":   "name1",
		"active": 1.0,
	}

	result, err = record.GetBool("active")
	if err != nil {
		t.Fatal(err)
	} else if !result {
		t.Logf("\nExpected %#v\nbut got  %#v\n", true, result)
		t.Fail()
	}
}

func Test_RecordGetFloat(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	result, err := record.GetFloat("value")
	if err != nil {
		t.Fatal(err)
	} else if result != 123.0 {
		t.Logf("\nExpected %f\nbut got  %#v\n", 123.0, result)
		t.Fail()
	}
}

func Test_RecordGetInt(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	result, err := record.GetInt("value")
	if err != nil {
		t.Fatal(err)
	} else if result != 123 {
		t.Logf("\nExpected %d\nbut got  %#v\n", 123, result)
		t.Fail()
	}
}

func Test_RecordGetSlice(t *testing.T) {
	record := Record{
		"name": "name1",
		"value": []interface{}{
			"value1",
			"value2",
		},
	}

	expected := []interface{}{"value1", "value2"}

	result, err := record.GetSlice("value")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %#v\nbut got  %#v\n", expected, result)
		t.Fail()
	}
}

func Test_RecordGetString(t *testing.T) {
	record := Record{
		"name":  "name1",
		"value": 123.0,
	}

	result, err := record.GetString("name")
	if err != nil {
		t.Fatal(err)
	} else if result != "name1" {
		t.Logf("\nExpected name1\nbut got  %#v\n", result)
		t.Fail()
	}
}

func Test_RecordGetTime(t *testing.T) {
	record := Record{
		"name": "name1",
		"time": 0.0,
	}

	expected := time.Unix(0, 0)

	result, err := record.GetTime("time")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %s\nbut got  %s\n", expected, result)
		t.Fail()
	}

	record = Record{
		"name": "name1",
		"time": 1.439633040e9,
	}

	expected = time.Unix(1439633040, 0)

	result, err = record.GetTime("time")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Logf("\nExpected %s\nbut got  %s\n", expected, result)
		t.Fail()
	}
}
