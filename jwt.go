package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"hash"
)

var (
	hasher     hash.Hash
	privateKey string
)

// Sets private key
func SetPrivateKey(key string) error {
	if key == "" {
		panic("Private Key is cannot Be Empty : Suggestion : use SetPrivatekey() function")
	}

	if privateKey != "" {
		panic("Private Key is Aldready Set , you Can't Use SetPrivateKey() Function Again")
	}

	privateKey = key
	hasher = hmac.New(sha256.New, []byte(privateKey))
	return nil
}

// Hashes STring to sha356 with private key set by function
// SetPrivateKey() and return error if private key is not set
func hashTheString(input string) (string, error) {
	if hasher == nil || privateKey == "" {
		return "", errors.New("set private key : suggestion : use SetPrivateKey() function")
	}

	hasher.Write([]byte(input))
	b := string(hasher.Sum(nil))
	hasher.Reset()

	return encode(string(b)), nil
}

// Function Which Encodes input String into base64 string
func encode(input string) (output string) {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Function Which Decodes input String into readable string
func decode(input string) (output string, err error) {
	b, err := base64.StdEncoding.DecodeString(input)
	return string(b), err
}
