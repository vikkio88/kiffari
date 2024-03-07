//go:build prod
// +build prod

package conf

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
)

const (
	LatestNotesLimit  = 10
	TrandingTagsLimit = 5
	Version           = "PROD_VERSION"
	GinMode           = "release"
	KiffariEnabled    = false
	BaseAddr          = ""
)

var Cors = []string{
	"https://vikkio.alwaysdata.net/",
}

var TokenSignature = os.Getenv("TOKEN_SIGNATURE")
var Port = os.Getenv("PORT")

var Connection = postgres.Open(fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s",
	os.Getenv("PG_HOST"),
	os.Getenv("PG_USER"),
	os.Getenv("PG_PASSWORD"),
	os.Getenv("PG_DBNAME"),
	os.Getenv("PG_PORT"),
))
