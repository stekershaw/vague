package vague

import (
	"encoding/json"
	"strconv"
)

// Int is a type aliasing 'int'
type Int int

// UnmarshalJSON : Int implements the json.Unmarshaler interface
// Internally an int64 is used (for consistency, as this is the return value
// from strvconv.ParseInt) but as the receiver is an int the caller should
// be aware of the possibility of overflow.
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	// try first to parse as an int
	var n int64
	if err = json.Unmarshal(b, &n); err == nil {
		*i = Int(n) // note this could overflow, calling code should check err.Err for this
		return
	}
	// if parsing as an int failed, try and parse an int from a string
	var s string
	if err = json.Unmarshal(b, &s); err == nil {
		if s == "" {
			*i = 0
			return
		}
		n, err = strconv.ParseInt(s, 10, 64)
		*i = Int(n)
	}
	return
}
