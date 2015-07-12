package jsons_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/bfontaine/jsons"
	"github.com/stretchr/testify/assert"
)

func TestReaderNextEmpty(t *testing.T) {
	b := bytes.NewReader([]byte{})
	r := jsons.NewReader(b)

	assert.NotNil(t, r)

	var m map[string]int

	assert.Equal(t, io.EOF, r.Next(&m))
}

func TestReaderNextWrongJSON(t *testing.T) {
	b := bytes.NewReader([]byte(`{"key":"`))
	r := jsons.NewReader(b)

	assert.NotNil(t, r)

	var m map[string]string

	assert.NotNil(t, r.Next(&m))
}

func TestReaderNext(t *testing.T) {
	b := bytes.NewReader([]byte("{\"k\":1}\n{\"k\":2,\"x\":11}\n"))
	r := jsons.NewReader(b)

	assert.NotNil(t, r)

	var m map[string]int

	assert.Nil(t, r.Next(&m))
	assert.Equal(t, 1, m["k"])

	assert.Nil(t, r.Next(&m))
	assert.Equal(t, 2, m["k"])
	assert.Equal(t, 11, m["x"])

	assert.Equal(t, io.EOF, r.Next(&m))
}
