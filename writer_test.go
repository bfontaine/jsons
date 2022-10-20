package jsons_test

import (
	"bytes"
	"testing"

	"github.com/bfontaine/jsons"
	"github.com/stretchr/testify/assert"
)

func TestWriterAdd(t *testing.T) {
	var b bytes.Buffer

	r := jsons.NewWriter(&b)
	assert.NotNil(t, r)

	assert.Nil(t, r.Add(map[string]int{
		"foo": 1,
	}))

	assert.Nil(t, r.AddAll(
		map[string]int{"foo": 2},
		map[string]bool{"foo": false}))

	assert.Nil(t, r.Add(struct {
		X int `json:"xyz"`
	}{X: 42}))

	s := b.String()

	assert.Equal(t, "{\"foo\":1}\n{\"foo\":2}\n{\"foo\":false}\n{\"xyz\":42}\n", s)
}
