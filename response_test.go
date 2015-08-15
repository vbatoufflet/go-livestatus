package livestatus

import (
	"testing"
)

func Test_ResponseLen(t *testing.T) {
	resp := Response{
		Status: 200,
		Records: []Record{
			Record{"name": "name1", "value": 123.0},
			Record{"name": "name2", "value": 456.0},
		},
	}

	if resp.Len() != 2 {
		t.Logf("\nExpected 2\nbut got  %#v\n", resp.Len())
		t.Fail()
	}
}
