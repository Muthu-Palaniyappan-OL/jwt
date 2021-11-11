package jwt

import (
	"errors"
	"net/http"
	"strings"
)

// There is a difference between authenticate and authorization
// authorization Means validating the user alone
// without getting all details.

// This Returns The Json String When is encoded and
// authenticated by the hash function
func authorize(rw http.ResponseWriter) (string, error) {
	jwt := rw.Header().Get("jwt")
	if jwt == "" {
		return "", errors.New("New Authentication Can't Authorise | jwt header is not set")
	}

	splitStrings := strings.Split(jwt, ".")

	if (splitStrings[1] != encode(splitStrings[0])){
		return "", nil
	}

	return splitStrings[0], nil
}