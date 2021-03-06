// Copyright 2021 Muthu Palaniyappan OL. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

package jwt

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorize(t *testing.T) {
	inputCookies := []string{
		"eyJOYW1lIjoidGVzdGluZ3ByaXZhdGVrZXkifQ==.kWPDUpVfpo4a5mrorXRqewieEpplX/MJZeWqutSmPVM=",
		"eyJOYW1lIjoibXV0aHUifQ==.zx8Xt8A03iBbeDRrkgGWJqDG9FV6zL5ExaK93XkQ8xc=",
	}

	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}
	for _, cookiesValue := range inputCookies {

		r := httptest.NewRequest("GET", "localhost:8080/", nil)
		r.AddCookie(&http.Cookie{
			Name:   "jwt",
			Value:  cookiesValue,
			MaxAge: 300,
		})
		s, err := Authorize(r)
		if err != nil {
			t.Fatal("error created by Authenticate() function, error:", err)
		}
	}
}
