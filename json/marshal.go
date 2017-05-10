package json

import (
	"encoding/json"
	"github.com/mono83/validate"
)

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
//
// Additionally, if v implements validate.Interface, applies validation checks
func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err == nil {
		if vv, ok := v.(validate.Interface); ok {
			err = vv.Validate()
		}
	}

	return err
}

// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
