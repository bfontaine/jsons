package jsons

import (
	"encoding/json"
	"io"
)

// Writer is a JSONS writer
type Writer struct {
	enc *json.Encoder
}

// NewWriter returns a new Writer, which writes JSONS-encoded data
// in the given io.Writer implementation.
func NewWriter(w io.Writer) Writer {
	return Writer{enc: json.NewEncoder(w)}
}

// Add encodes the given value and write it as a JSON object.
func (jw Writer) Add(v interface{}) error {
	return jw.enc.Encode(v)
}

// AddAll is equivalent to calling Add on each of its arguments
func (jw Writer) AddAll(args ...interface{}) (err error) {
	for _, v := range args {
		if err = jw.Add(v); err != nil {
			return
		}
	}
	return
}
