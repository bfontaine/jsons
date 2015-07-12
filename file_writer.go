package jsons

import (
	"io"
	"os"
)

type FileWriter struct {
	filename string

	f io.WriteCloser
	w *Writer
}

func NewFileWriter(filename string) *FileWriter {
	return &FileWriter{filename: filename}
}

func (fw *FileWriter) Open() error {
	f, err := os.Create(fw.filename)
	if err != nil {
		return err
	}
	fw.f = f
	fw.w = NewWriter(f)
	return nil
}

func (fw *FileWriter) Close() error {
	return fw.f.Close()
}

func (fw *FileWriter) Add(v interface{}) error {
	return fw.w.Add(v)
}
