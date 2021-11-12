// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

/*
	package jwt implements simple way to create jwt tokens for your web project in golang.

	This package takes HMAC algorithm by default for any other algorithm download this source code and change it freely.

	Authenticate means putting username and password and getting verified and generate key.
 	Authorization means having a key which can be used to open door.

 	Functions implemented and exported in this package:

 	func Authenticate(rw http.ResponseWriter, jsonStringTOSend string, seconds int)
 	func Authorize(r *http.Request) (string, error)
 	func SetCookiePath(path string)
 	func IsPrivateKeySet() bool

*/
package jwt
