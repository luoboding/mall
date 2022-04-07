package utils

import (
	"crypto/sha256"

	"fmt"
)

const (
	SALT = "mall"
)

func If(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func Encrypt(input string) string {
	sum := sha256.Sum256([]byte(input))
	result := sha256.Sum256([]byte(fmt.Sprintf("%x", sum) + SALT))
	return fmt.Sprintf("%x", result)
}
