package jwt

import (
	"math/rand"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}

	rw := httptest.NewRecorder()

	err := Authenticate(rw, encode("testingjsonstring"), rand.Int())

	if err != nil {
		t.Log("Error Created By Func ", err)
		return
	}

	cookies := rw.Result().Cookies()
	jwtExists := false

	for _, c := range cookies {
		if c.Name == "jwt" {
			jwtExists = true
		}
	}

	if jwtExists == true {

	} else {
		t.Fatal("jwt Cookies in not Set in resposnse")
	}

}
