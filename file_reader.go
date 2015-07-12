package jsons

import (
	"io"
	"os"
)

type FileReader struct {
	filename string
	f        io.WriteCloser
	r        *Reader
}

func NewFileReader(filename string) *FileReader {
	return &FileReader{filename: filename}
}

func (fr *FileReader) Open() error {
	f, err := os.Open(fr.filename)
	if err != nil {
		return err
	}
	fr.f = f
	fr.r = NewReader(f)
	return nil
}

func (fr *FileReader) Close() error {
	return fr.f.Close()
}

func (fr *FileReader) Next(v interface{}) error {
	return fr.r.Next(v)
}
