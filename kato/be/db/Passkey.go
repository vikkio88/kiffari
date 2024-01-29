package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PasskeyClear struct {
	Key string `json:"passkey" binding:"required"`
}

func (p *PasskeyClear) Check(key Passkey) bool {
	return bcrypt.CompareHashAndPassword([]byte(key.Hash), []byte(p.Key)) == nil
}

type Passkey struct {
	gorm.Model
	Hash string
}

func NewPasskey(clear string) Passkey {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(clear), bcrypt.DefaultCost)
	return Passkey{Hash: string(hashed)}
}
