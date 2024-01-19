package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct {
	g *gorm.DB
}

func NewDb(fileName string) *Db {
	g, err := gorm.Open(
		sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", fileName)),
		&gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	migrate(g)
	return &Db{g}
}
