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
func Authorize(r *http.Request) (string, error) {

	var jwt string

	cs := r.Cookies()

	for _, cookie := range cs {
		if cookie.Name == "jwt" {
			jwt = cookie.Value
		}
	}

	if jwt == "" {
		return "", errors.New("jwt cookie is not set or empty")
	}

	splitStrings := strings.Split(jwt, ".")

	if len(splitStrings) != 2 {
		return splitStrings[0], errors.New("jwt doesnt have exactly one dot separating")
	}

	str, err := hashTheString(splitStrings[0])

	if err != nil {
		return "", err
	}

	if splitStrings[1] != str {
		return splitStrings[0] + "/\\" + str + "|||" + splitStrings[1], errors.New("signature matching failed")
	}

	decodedString, err := decode(splitStrings[0])
	if err != nil {
		return "", err
	}

	return decodedString, nil
}
