// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

package jwt

import (
	"net/http"
)

var jwtCookie *http.Cookie = &http.Cookie{
	Name:     "jwt",
	SameSite: http.SameSiteDefaultMode,
	Secure:   true,
}

// After Authenticating the user by verifying the username in database you need to create a json string by using json.Marshal() Function and send it via this function which gets converted into encoded hash and send it user via http.ResponseWriter as a token you could specify duration of cookie in seconds argument.
func Authenticate(rw http.ResponseWriter, jsonStringTOSend string, seconds int) error {
	s := encode(jsonStringTOSend)
	h, err := hashTheString(s)
	if err != nil {
		return err
	}
	jwtCookie.Value = s + "." + h
	jwtCookie.MaxAge = seconds
	http.SetCookie(rw, jwtCookie)
	return nil
}

// Set jwt cookie's path using this function
func SetCookiePath(path string) {
	if path == "" {
		return
	}
	jwtCookie.Path = path
}
