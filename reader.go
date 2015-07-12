package jsons

import (
	"encoding/json"
	"io"
)

// Reader is a JSONS reader
type Reader struct {
	dec *json.Decoder
}

// NewReader returns a pointer on a new Reader which consumes its data from the
// given io.Reader implementation.
func NewReader(r io.Reader) *Reader {
	return &Reader{dec: json.NewDecoder(r)}
}

// Next unmarshal the next JSON object in the given value, which must be a
// pointer.
func (r *Reader) Next(v interface{}) error {
	return r.dec.Decode(v)
}
