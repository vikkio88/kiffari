package libs

import "golang.org/x/crypto/bcrypt"

func Hash(clear string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(clear), bcrypt.DefaultCost)

	return string(bytes)
}

func CompareHash(hash string, clear string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(clear)) == nil
}
