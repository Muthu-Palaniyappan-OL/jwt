package jwt

import (
	"net/http"
)

// There is a difference between authenticate and authorization
// Authenticate Means Take in Username and password and
// Validatin the user

// Use This funtion when user is authorised after getting
// validation from database
func Authenticate(rw http.ResponseWriter, jsonStringTOSend string, seconds int) error {
	s := encode(jsonStringTOSend)
	h, err := hashTheString(s)
	if err != nil {
		return err
	}
	http.SetCookie(rw, &http.Cookie{
		Name:   "jwt",
		Value:  s + "." + h,
		MaxAge: seconds,
	})
	return nil
}
