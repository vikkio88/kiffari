package db

import (
	"errors"
	"kato-be/conf"
	"kato-be/libs"

	"gorm.io/gorm"
)

type Db struct {
	g *gorm.DB

	pk *Passkey
}

func NewDb() *Db {
	g, err := gorm.Open(
		conf.Connection,
		&gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	migrate(g)
	db := &Db{g, nil}
	return db
}

func (db *Db) IsTokenValid(token string) bool {
	if token == "" {
		return false
	}
	return libs.VerifyToken(token, conf.TokenSignature)
}

func (db *Db) CheckPk(pk PasskeyClear) (string, error) {
	pkd, isInit := db.getPasskey()
	if !isInit {
		return "", errors.New("not initialised")
	}

	if !pk.Check(pkd) {
		return "", errors.New("invalid passkey")
	}

	return libs.NewToken(conf.TokenSignature)
}

func (db *Db) getPasskey() (*Passkey, bool) {
	if db.pk == nil {
		p, ok := db.GetHashedKey()
		if ok {
			db.pk = p
		}
		return p, ok
	}
	return db.pk, true
}
