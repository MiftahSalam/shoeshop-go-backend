package middleware

import (
	"github.com/google/uuid"
)

// GenerateUUID produces random ID based on UUID
func GenerateUUID() (string, error) {
	_uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return _uuid.String(), nil
}
