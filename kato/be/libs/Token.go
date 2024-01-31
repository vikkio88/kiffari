package libs

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(passkey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	value, err := token.SignedString([]byte(passkey))

	if err != nil {
		return "", err
	}

	return value, nil
}

func VerifyToken(token string, passkey string) bool {
	value, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(passkey), nil
	})

	fmt.Println("valid?", value.Valid)

	if err != nil || !value.Valid {
		return false
	}

	exp, _ := value.Claims.GetExpirationTime()
	fmt.Println("exp", exp.Unix())
	return exp.After(time.Now())
}
