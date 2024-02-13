//go:build !prod
// +build !prod

package conf

const (
	TokenSignature   = "someTokensignature"
	LatestNotesLimit = 5
	Version          = "DEV"
	GinMode          = "debug"
)
