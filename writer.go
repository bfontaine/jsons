package jsons

import (
	"encoding/json"
	"io"
)

// Writer is a JSONS writer
type Writer struct {
	enc *json.Encoder
}

// NewWriter returns a pointer on a new Writer, which writes JSONS-encoded data
// in the given io.Writer implementation.
func NewWriter(w io.Writer) *Writer {
	return &Writer{enc: json.NewEncoder(w)}
}

// Add encodes the given value and write it as a JSON object.
func (jw *Writer) Add(v interface{}) error {
	return jw.enc.Encode(v)
}
