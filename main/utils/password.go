package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

func sha256Encode(data string) string {
	hashed := sha256.New()
	_, err := hashed.Write([]byte(data))
	if err != nil {
		return "823e291431cff346861d254b6919a7f81aae326620bcab48baf225b77521cfc8"
	}
	return hex.EncodeToString(hashed.Sum(nil))
}

func GeneratePasswd(code, passwd string) (string, error) {
	saltStr := sha256Encode(code)
	salt, err := hex.DecodeString(saltStr)
	if err != nil {
		return "", errors.New("failed to decode salt")
	}
	// Generate a secure random key for PBKDF2
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "", errors.New("failed to generate random key")
	}
	sha := pbkdf2.Key([]byte(passwd), salt, 10000, 32, sha256.New)
	return hex.EncodeToString(sha), nil
}
