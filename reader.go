package jsons

import (
	"encoding/json"
	"io"
)

type Reader struct {
	dec *json.Decoder
}

func NewReader(r io.Reader) *Reader {
	return &Reader{dec: json.NewDecoder(r)}
}

func (r *Reader) Next(v interface{}) error {
	return r.dec.Decode(v)
}
