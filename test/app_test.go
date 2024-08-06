package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	word := "hello"
	assert.NotNil(t, word)
	assert.Equal(t, "hello", word)
}
