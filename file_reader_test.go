package jsons_test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bfontaine/jsons"
	"github.com/bfontaine/vanish/vanish"
	"github.com/stretchr/testify/assert"
)

func TestFileReaderOpenNonExistingFile(t *testing.T) {
	vanish.File(func(name string) {
		assert.Nil(t, os.Remove(name))

		fr := jsons.NewFileReader(name)

		assert.NotNil(t, fr.Open())
	})
}

func TestFileReaderEmptyFile(t *testing.T) {
	vanish.File(func(name string) {
		fr := jsons.NewFileReader(name)

		assert.Nil(t, fr.Open())

		var m map[int]int

		assert.Equal(t, io.EOF, fr.Next(&m))
		assert.Nil(t, fr.Close())
	})
}

func TestFileReaderNext(t *testing.T) {
	vanish.File(func(name string) {

		ioutil.WriteFile(name, []byte("{\"foo\": 17}\n"), 0644)

		fr := jsons.NewFileReader(name)

		assert.Nil(t, fr.Open())

		var m map[string]int

		assert.Nil(t, fr.Next(&m))
		assert.Equal(t, 17, m["foo"])

		assert.Equal(t, io.EOF, fr.Next(&m))
		assert.Nil(t, fr.Close())
	})
}
