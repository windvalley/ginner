package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

// EncryptByPublic encrypt str by public key
func EncryptByPublic(src, path string) (string, error) {
	block, err := getPemBlock(path)
	if err != nil {
		return "", err
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	signatureBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(src))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(signatureBytes), nil
}

// DecryptByPrivate decrypt signature by private key
func DecryptByPrivate(signature, path string) (string, error) {
	block, err := getPemBlock(path)
	if err != nil {
		return "", err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	srcBytes, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return "", err
	}

	srcBytes, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, srcBytes)
	if err != nil {
		return "", err
	}

	return string(srcBytes), nil
}

func getPemBlock(path string) (*pem.Block, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	return block, nil
}
