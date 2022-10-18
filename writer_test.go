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
		"foo": 42,
	}))

	assert.Nil(t, r.Add(struct {
		X int `json:"xyz"`
	}{X: 42}))

	s := b.String()

	assert.Equal(t, "{\"foo\":42}\n{\"xyz\":42}\n", s)
}
