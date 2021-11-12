// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

package jwt

import (
	"errors"
	"net/http"
	"strings"
)

// reads jwt cookies in the client http.Request and checks the validity of jwt cookie signature (check whether it is signed with the same privateKey) and returns the json string which was set as a cookie by server before in Authenticate() function.
func Authorize(r *http.Request) (string, error) {
	if !IsPrivateKeySet() {
		panic("private key not set")
	}

	var jwt string
	cs := r.Cookies()
	for _, cookie := range cs {
		if cookie.Name == "jwt" {
			jwt = cookie.Value
		}
	}
	if jwt == "" {
		return "", errors.New("jwt cookie is not set or jwt cookie is empty")
	}

	splitStrings := strings.Split(jwt, ".")
	if len(splitStrings) != 2 {
		return splitStrings[0], errors.New("jwt token is tampered")
	}

	str, err := hashTheString(splitStrings[0])
	if err != nil {
		return "", err
	}
	if splitStrings[1] != str {
		return "", errors.New("signature does not match")
	}
	decodedString, err := decode(splitStrings[0])
	if err != nil {
		return "", err
	}

	return decodedString, nil
}
