// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

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

	_, err := Authenticate(rw, encode("testingjsonstring"), rand.Int())

	if err != nil {
		t.Fatalf("error created by Authenticate() function, error:", err)
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
		t.Fatal("jwt Cookies is not set")
	}

}
