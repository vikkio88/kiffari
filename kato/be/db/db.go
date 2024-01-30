package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (db *Db) CheckPk(pk PasskeyClear) (string, error) {
	pkd := db.getPasskey()
	if !pk.Check(pkd) {
		return "", errors.New("Invalid Passkey")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(pk.Key))
}

func (db *Db) getPasskey() Passkey {
	return NewPasskey("password")
}
