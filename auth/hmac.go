package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"hash"
)

// Hmac HMAC signature
func Hmac(deaName, src, key string) (string, error) {
	var hm hash.Hash

	switch deaName {
	case "hmac_sha256":
		hm = hmac.New(sha256.New, []byte(key))
	case "hmac_sha1":
		hm = hmac.New(sha1.New, []byte(key))
	case "hmac_md5":
		hm = hmac.New(md5.New, []byte(key))
	default:
		return "", errors.New("unknown DEA")
	}

	hm.Write([]byte(src))

	signatureBytes := hm.Sum(nil)

	signature := base64.URLEncoding.EncodeToString(signatureBytes)

	return signature, nil
}
