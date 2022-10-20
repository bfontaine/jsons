package jsons

import (
	"io"
	"os"
)

// FileReader is a wrapper around a Reader to read JSONS data from a file.
type FileReader struct {
	filename string
	f        io.WriteCloser
	r        *Reader
}

// NewFileReader returns a pointer on a new FileReader which will read its data
// from the given file.
func NewFileReader(filename string) *FileReader {
	return &FileReader{filename: filename}
}

// Open opens the underlying file for reading. You must call this before any
// Close or Next call.
func (fr *FileReader) Open() error {
	f, err := os.Open(fr.filename)
	if err != nil {
		return err
	}
	fr.f = f
	fr.r = NewReader(f)
	return nil
}

// Close closes the underlying file.
func (fr *FileReader) Close() error {
	fr.r = nil
	err := fr.f.Close()
	fr.f = nil
	return err
}

// Next unmarshals the next JSON object in the given value, which must be a
// pointer.
func (fr *FileReader) Next(v interface{}) error {
	return fr.r.Next(v)
}
