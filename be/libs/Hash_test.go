package libs_test

import (
	"kato-be/libs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	clear := "password"
	assert.NotEqual(t, clear, libs.Hash(clear))
}

func TestCompare(t *testing.T) {
	clear := "password"

	hash := libs.Hash(clear)

	assert.True(t, libs.CompareHash(hash, clear))
	assert.False(t, libs.CompareHash(hash, "not password"))
}
