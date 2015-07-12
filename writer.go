package jsons

import (
	"encoding/json"
	"io"
)

type Writer struct {
	enc *json.Encoder
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{enc: json.NewEncoder(w)}
}

func (jw *Writer) Add(v interface{}) error {
	return jw.enc.Encode(v)
}
