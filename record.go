package livestatus

import (
	"reflect"
	"sort"
	"time"
)

// Record represents a Livestatus response entry.
type Record map[string]interface{}

// Len returns the number of columns present in the record.
func (r Record) Len() int {
	return len(r)
}

// Columns returns the list of the record columns.
func (r Record) Columns() []string {
	cols := []string{}
	for k := range r {
		cols = append(cols, k)
	}
	sort.Strings(cols)

	return cols
}

// Get returns an interface value for a specific column.
func (r Record) Get(column string) (interface{}, error) {
	v, ok := r[column]
	if !ok {
		return nil, ErrUnknownColumn
	}

	return v, nil
}

// GetBool returns a boolean value for a specific column.
func (r Record) GetBool(column string) (bool, error) {
	v, err := r.getKey(reflect.Float64, column)
	if err != nil {
		return false, err
	}

	return v == 1.0, nil
}

// GetFloat returns a float value for a specific column.
func (r Record) GetFloat(column string) (float64, error) {
	v, err := r.getKey(reflect.Float64, column)
	if err != nil {
		return 0, err
	}

	return v.(float64), nil
}

// GetInt returns an integer value for a specific column.
func (r Record) GetInt(column string) (int64, error) {
	v, err := r.getKey(reflect.Float64, column)
	if err != nil {
		return 0, err
	}

	return int64(v.(float64)), nil
}

// GetSlice returns a slice of interface value for a specific column.
func (r Record) GetSlice(column string) ([]interface{}, error) {
	v, err := r.getKey(reflect.Slice, column)
	if err != nil {
		return nil, err
	}

	rv := reflect.ValueOf(v)
	n := rv.Len()
	out := make([]interface{}, n)

	for i := 0; i < n; i++ {
		out[i] = rv.Index(i).Interface()
	}

	return out, nil
}

// GetString returns a string value for a specific column.
func (r Record) GetString(column string) (string, error) {
	v, err := r.getKey(reflect.String, column)
	if err != nil {
		return "", err
	}

	return v.(string), nil
}

// GetTime returns a time struct value for a specific column.
func (r Record) GetTime(column string) (time.Time, error) {
	v, err := r.getKey(reflect.Float64, column)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(int64(v.(float64)), 0), nil
}

func (r Record) getKey(k reflect.Kind, column string) (interface{}, error) {
	val, ok := r[column]
	if !ok {
		return nil, ErrUnknownColumn
	} else if vk := reflect.ValueOf(val).Kind(); vk != k {
		return nil, ErrInvalidType
	}

	return val, nil
}
