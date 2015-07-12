package jsons

import "os"

type FileReader struct {
	Err error

	filename string
}

func NewFileReader(filename string) *FileReader {
	return &FileReader{filename: filename}
}

func (fr *FileReader) Chan() (chan interface{}, error) {
	f, err := os.Open(fr.filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := NewReader(f)

	c := make(chan interface{})

	go func() {
		var el interface{}

		if err := reader.Next(&el); err != nil {
			fr.Err = err
			close(c)
			return
		}

		c <- el
	}()

	return c, nil
}
