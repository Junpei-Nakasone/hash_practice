package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHash は文字列をハッシュします。
func GenerateHash(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Verify(hash string, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	if err != nil {
		return false
	}
	return true
}

func main() {
	test, _ := GenerateHash("password")

	verified := Verify(test, "password")
	fmt.Println(verified)
}
