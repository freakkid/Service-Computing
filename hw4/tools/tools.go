package tools

import (
	"crypto/md5"
	"fmt"
	"io"

	uuid "github.com/satori/go.uuid"
)

// simple tools used in this program
// 1. encrypt password by MD5
// 2. creating UUID as id of user

// to encrypt password by MD5
func MD5Encryption(text string) string {
	hash := md5.New()
	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// creating UUID as id of user
// return string
func GetUUID() string {
	return uuid.NewV4().String()
}
