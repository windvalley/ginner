package util

import (
	uuid "github.com/satori/go.uuid"
)

// GenerateUUID generate uuid
func GenerateUUID() (string, error) {
	u4, err := uuid.NewV4()
	return u4.String(), err
}
