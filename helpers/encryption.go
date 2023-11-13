package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type EncryptionHelper struct{}

func (h *EncryptionHelper) GenerateRandomHex(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), errors.New("canonroeroejo")
}

func (h *EncryptionHelper) SHA256(input string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
