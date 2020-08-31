package auth

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5sum get value of md5sum
func Md5sum(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
