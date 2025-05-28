package pkg

import (
	crand "crypto/rand"
	"encoding/hex"
	mathrand "math/rand"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func HashSecret(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	return string(bytes), err
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[mathrand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := crand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
