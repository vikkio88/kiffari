//go:build !prod
// +build !prod

package conf

import (
	"gorm.io/driver/postgres"
	// "gorm.io/driver/sqlite"
)

const (
	TokenSignature   = "someTokensignature"
	LatestNotesLimit = 5
	Version          = "DEV"
	GinMode          = "debug"
)

// var Connection = sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", "testing.db"))
var Connection = postgres.Open("host=localhost user=kato password=kato dbname=kato port=5432")
