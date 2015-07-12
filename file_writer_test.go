package jsons_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bfontaine/jsons"
	"github.com/bfontaine/vanish/vanish"
	"github.com/stretchr/testify/assert"
)

func TestFileWriterCreateFileIfItDoesntExist(t *testing.T) {
	vanish.File(func(name string) {
		assert.Nil(t, os.Remove(name))

		fw := jsons.NewFileWriter(name)
		assert.NotNil(t, fw)

		assert.Nil(t, fw.Open())
		assert.Nil(t, fw.Close())

		_, err := os.Stat(name)
		assert.Nil(t, err)
	})
}

func TestFileWriterAdd(t *testing.T) {
	vanish.File(func(name string) {
		fw := jsons.NewFileWriter(name)
		assert.NotNil(t, fw)

		assert.Nil(t, fw.Open())

		assert.Nil(t, fw.Add(map[string]string{"foo": "bar"}))
		assert.Nil(t, fw.Add(struct{ A float64 }{A: 42.3}))

		assert.Nil(t, fw.Close())

		content, err := ioutil.ReadFile(name)
		assert.Nil(t, err)

		assert.Equal(t, "{\"foo\":\"bar\"}\n{\"A\":42.3}\n", string(content))
	})
}
