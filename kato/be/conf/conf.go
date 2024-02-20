//go:build !prod
// +build !prod

package conf

import (
	"fmt"

	"gorm.io/driver/sqlite"
)

const (
	TokenSignature   = "someTokensignature"
	LatestNotesLimit = 5
	Version          = "DEV"
	GinMode          = "debug"
	Port             = "5111"
)

var Cors = []string{
	"http://localhost:5111",
	"http://127.0.0.1:5111",
	"http://localhost:5173",
	"http://127.0.0.1:5173",
}

// If you need to connect on the docker instance
// var Connection = postgres.Open("host=localhost user=kato password=kato dbname=kato port=5432")
var Connection = sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", "testing.db"))
