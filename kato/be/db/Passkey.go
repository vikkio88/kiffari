package db

import (
	"kato-be/libs"

	"gorm.io/gorm"
)

type PasskeyClear struct {
	Key string `json:"passkey" binding:"required"`
}

func (p *PasskeyClear) Check(key Passkey) bool {
	return libs.CompareHash(key.Hash, p.Key)
}

type Passkey struct {
	gorm.Model
	Hash string
}

func NewPasskey(clear string) Passkey {
	return Passkey{Hash: libs.Hash(clear)}
}
