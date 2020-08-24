package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// aes-128-cbc
func AESEncrypt(src, key string) (string, error) {
	srcBytes, keyBytes := []byte(src), []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	srcBytes = PKCS7Padding(srcBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])

	signatureBytes := make([]byte, len(srcBytes))
	blockMode.CryptBlocks(signatureBytes, srcBytes)
	signature := base64.URLEncoding.EncodeToString(signatureBytes)

	return signature, nil
}

func AESDecrypt(signature, key string) (string, error) {
	keyBytes := []byte(key)
	signatureBytes, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])

	srcBytes := make([]byte, len(signatureBytes))
	blockMode.CryptBlocks(srcBytes, signatureBytes)
	srcBytes = PKCS7UnPadding(srcBytes)

	return string(srcBytes), nil
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:length-unpadding]
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padtext...)
}
