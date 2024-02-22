package libs_test

import (
	"kato-be/libs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenBuilder(t *testing.T) {
	password := "password"
	token, _ := libs.NewToken(password)

	assert.True(t, libs.VerifyToken(token, password))

	assert.False(t, libs.VerifyToken(token, "notpassword"))
}
