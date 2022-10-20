package jsons

import (
	"io"
	"os"
)

// FileWriter is a wrapper around a Writer which writes JSONS data in a file.
type FileWriter struct {
	filename string

	f io.WriteCloser
	w Writer
}

// NewFileWriter returns a pointer on a new FileWriter which will write in the
// given filename. The file is truncated if it already exists.
func NewFileWriter(filename string) *FileWriter {
	// TODO support append?
	return &FileWriter{filename: filename}
}

// Open opens the underlying file for writing. This must be called before any
// Add or Close call.
func (fw *FileWriter) Open() error {
	f, err := os.Create(fw.filename)
	if err != nil {
		return err
	}
	fw.f = f
	fw.w = NewWriter(f)
	return nil
}

// Close closes the underlying file. You must call this when you're done adding
// values.
func (fw *FileWriter) Close() error {
	return fw.f.Close()
}

// Add encodes the given value as a JSON object and write it in the underlying
// file.
func (fw *FileWriter) Add(v interface{}) error {
	return fw.w.Add(v)
}
