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
func Authorize(r http.Response) (string, error) {
	var jwt string
	cs := r.Cookies()
	for _, cookie := range cs {
		if cookie.Name == "jwt" {
			jwt = cookie.Name
		}
	}
	jwt, err := decode(jwt)
	if err != nil {
		return "", errors.New("connat decode")
	}
	if jwt == "" {
		return "", errors.New("or not cookies is set | new authentication can't authorise | jwt header is not set")
	}

	splitStrings := strings.Split(jwt, ".")

	if splitStrings[1] != encode(splitStrings[0]) {
		return "", nil
	}

	return splitStrings[0], nil
}
