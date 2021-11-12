package jwt

import (
	"net/http/httptest"
	"os"
	"testing"
	"net/http"
)

func TestAuthorize(t *testing.T){
	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}
	r := httptest.NewRequest("GET", "localhost:8080/", nil)
	r.AddCookie(&http.Cookie{
		Name: "jwt",
		Value: "eyJOYW1lIjoiTXV0aHUifQ==.HGKHCBakRfMlBU69kjRsJmFbVlF1p0bx0RDW8JgojlA=",
		MaxAge: 300,
	})
	s, err := Authorize(r)
	if err != nil {
		t.Fatal("Error From Authorize() Function:", err, "Returned String => ", s)
	}
	t.Log(r.Write(os.Stdout))
	t.Log("String:", s)
}

func TestSimple(t *testing.T){
	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}
	t.Log(hashTheString("eyJOYW1lIjoiTXV0aHUifQ=="))
}