package db

import (
	"errors"
	"fmt"
	"kato-be/conf"
	"kato-be/libs"

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

func (db *Db) IsTokenValid(token string) bool {
	if token == "" {
		return false
	}
	return libs.VerifyToken(token, conf.TOKEN_SIGNATURE)
}

func (db *Db) CheckPk(pk PasskeyClear) (string, error) {
	pkd := db.getPasskey()
	if !pk.Check(pkd) {
		return "", errors.New("invalid passkey")
	}

	return libs.NewToken(conf.TOKEN_SIGNATURE)
}

func (db *Db) getPasskey() Passkey {
	//TODO: cache this
	return NewPasskey("password")
}
