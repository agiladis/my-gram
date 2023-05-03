package helper

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 10
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(plainPass, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass))
}
