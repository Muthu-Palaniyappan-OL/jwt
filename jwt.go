// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

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

// returns true is pivateKey is set
func IsPrivateKeySet() bool {
	if privateKey == "" {
		return false
	} else {
		return true
	}
}

// sets given string as a privateKey. privateKey can be set only once.
func SetPrivateKey(key string) error {
	if key == "" {
		panic("Empty Private Key Can't Be Set")
	}

	if privateKey != "" {
		panic("Private Key is Aldready Set , you Can't Use SetPrivateKey() Function Again")
	}

	privateKey = key
	hasher = hmac.New(sha256.New, []byte(privateKey))
	return nil
}

// Hashes given String to SHA256 with privateKey set by function SetPrivateKey() and return error if private key is not set
func hashTheString(input string) (string, error) {
	if hasher == nil || !IsPrivateKeySet() {
		return "", errors.New("privateKey not set")
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
