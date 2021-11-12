# jwt

Simple yet Powerful Json Webtoken Library. Add JSON WebToken Authorisation in your web project in few clicks and codes.  package `jwt` implements jwt in a different way it remove overhead of calculation which hashing algorithm to use by takein `SHA256` by default agorithm and reduces JWT length size by removing header.

---

* [About](#about)
* [Install](#install)
* [Examples](#examples)

---

## About

jwt package takes SHA256 as default algorithm. This package doesn't implement Standard json web token. It uses a compressed version of jwt (compressed by author). the client could potentionall save few bytes of data by using compressed version. Don't store critical data as json webtoken as jwt does only base64 encryption.

## Install

```sh
go get -u github.com/Muthu-Palaniyappan-OL/jwt
```

## Examples

```go
func main(){
	jwt.SetPrivateKey("drjyt347yu3h87yuh")
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){

			// validate form data of the user

			jwt.Authenticate(rw, "json-string", 30000)

			// After setting token then start writing response

			// Start writing response
			rw.WriteHeader(200)
		})
	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request){
			json, err := jwt.Authenticate(rw, "json-string", 30000)

			// json is the string which you assiged the client when he first used GET /login request
		})
	http.ListenAndServe(":8080")
}
```