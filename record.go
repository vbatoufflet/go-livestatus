package livestatus

import (
	"sort"
	"time"
)

// Record is query response entry.
type Record map[string]interface{}

// Len returns the number of columns present in the record.
func (r Record) Len() int {
	return len(r)
}

// Columns returns the names of the columns present in the record.
func (r Record) Columns() []string {
	var cols []string

	for k := range r {
		cols = append(cols, k)
	}
	sort.Strings(cols)

	return cols
}

// Get returns an interface value for a specific column
//
// Returns an  error if the column is unknown.
func (r Record) Get(name string) (interface{}, error) {
	if v, ok := r[name]; ok {
		return v, nil
	}
	return nil, ErrUnknownColumn
}

// GetBool returns a boolean value for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as a boolean.
func (r Record) GetBool(name string) (bool, error) {
	v, err := r.Get(name)
	if err != nil {
		return false, err
	}
	vc, ok := v.(float64)
	if !ok {
		return false, ErrInvalidValue
	}
	return vc == 1, nil
}

// GetFloat returns a float value for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as a float.
func (r Record) GetFloat(name string) (float64, error) {
	v, err := r.Get(name)
	if err != nil {
		return 0, err
	}
	vc, ok := v.(float64)
	if !ok {
		return 0, ErrInvalidValue
	}
	return vc, nil
}

// GetInt returns an integer value for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as an integer.
func (r Record) GetInt(name string) (int64, error) {
	v, err := r.Get(name)
	if err != nil {
		return 0, err
	}
	vc, ok := v.(float64)
	if !ok {
		return 0, ErrInvalidValue
	}
	return int64(vc), nil
}

// GetSlice returns a slice of interfaces for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as a slice.
func (r Record) GetSlice(name string) ([]interface{}, error) {
	v, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	vc, ok := v.([]interface{})
	if !ok {
		return nil, ErrInvalidValue
	}
	return vc, nil
}

// GetString returns a string value for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as a string.
func (r Record) GetString(name string) (string, error) {
	v, err := r.Get(name)
	if err != nil {
		return "", err
	}
	vc, ok := v.(string)
	if !ok {
		return "", ErrInvalidValue
	}
	return vc, nil
}

// GetTime returns a time struct for a specific column.
//
// Returns an error if the column is unknown or if the value can't be represented as a time struct.
func (r Record) GetTime(name string) (time.Time, error) {
	v, err := r.Get(name)
	if err != nil {
		return time.Time{}, err
	}
	vc, ok := v.(float64)
	if !ok {
		return time.Time{}, ErrInvalidValue
	}
	return time.Unix(int64(vc), 0), nil
}

func (r Record) set(name string, v interface{}) {
	r[name] = v
}
