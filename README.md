# jwt
[![Go Reference](https://pkg.go.dev/badge/github.com/Muthu-Palaniyappan-OL/jwt.svg)](https://pkg.go.dev/github.com/Muthu-Palaniyappan-OL/jwt)

Simple yet Powerful Json Webtoken Library. Add JSON WebToken Authorisation in your web project in few clicks and codes.  package `jwt` implements jwt in a different way it remove overhead of calculation which hashing algorithm to use by takein `SHA256` by default agorithm and reduces JWT length size by removing header.

---

* [About](#about)
* [Install](#install)
* [Example](#example)

---

## About

__jwt__ package takes SHA256 as default algorithm. This package doesn't implement Standard json web token. It uses a compressed version of jwt (compressed by author). the client could potentionaly save few bytes of data by using this compressed version. Don't store your client's critical data as any json webtoken as jwt does only base64 encryption.

## Install

```sh
go get -u github.com/Muthu-Palaniyappan-OL/jwt
```

## Example

### Sample Code
```go
package main

import (
	"log"
	"net/http"

	"github.com/Muthu-Palaniyappan-OL/jwt"
)

func main() {
	jwt.SetPrivateKey("drjyt347yu3h87yuh")
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {

		// validate form data of the user

		jwt.Authenticate(rw, "json-string", 30000)
		rw.WriteHeader(200) // run jwt.Authenticate before writing into wr.WriteHeader()

		// After setting token then start writing response

		// Start writing response
	})
	http.HandleFunc("/feed", func(rw http.ResponseWriter, r *http.Request) {
		json, err := jwt.Authorize(r)

		// json is the string which you assiged the client when he first used GET /login request

		if err != nil {
			log.Println(err)
		}

		rw.Write([]byte(json))
	})
	http.ListenAndServe(":8080", nil)
}
```
