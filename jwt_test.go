package jwt

import (
	"testing"
)

func TestPrivatehash(t *testing.T) {

	hashs := []struct {
		Unhashed string
		Hashed   string
	}{
		{"eyJOYW1lIjoidGVzdGluZ3ByaXZhdGVrZXkifQ==", "kWPDUpVfpo4a5mrorXRqewieEpplX/MJZeWqutSmPVM="},
		{"eyJOYW1lIjoibXV0aHUifQ==", "zx8Xt8A03iBbeDRrkgGWJqDG9FV6zL5ExaK93XkQ8xc="},
	}

	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}

	for _, hash := range hashs {
		h, err := hashTheString(hash.Unhashed)
		if err != nil {
			t.Fatal(err)
		}
		if h != hash.Hashed {
			t.Fatalf("For Input %s Expected output %s recived output %s", hash.Unhashed, hash.Hashed, h)
		}
	}

}

func TestPrivateencode(t *testing.T) {

	encodingTests := []struct {
		UnEncodedString string
		EncodedString   string
	}{
		{"golang", "Z29sYW5n"},
		{"muthu", "bXV0aHU="},
		{"test", "dGVzdA=="},
	}

	for _, test := range encodingTests {
		if encode(test.UnEncodedString) != test.EncodedString {
			t.Logf("Expected encode(%s)=%s But Received %s", test.UnEncodedString, test.EncodedString, encode(test.UnEncodedString))
			t.Fatal("Not Passing encode() func tests")
		}
	}

}

func TestPrivatedecode(t *testing.T) {

	encodingTests := []struct {
		UnEncodedString string
		EncodedString   string
	}{
		{"golang", "Z29sYW5n"},
		{"muthu", "bXV0aHU="},
		{"test", "dGVzdA=="},
	}

	for _, test := range encodingTests {
		s, err := decode(test.EncodedString)
		if s != test.UnEncodedString || err != nil {
			t.Logf("Expected decode(%s)=%s But Received %s", test.EncodedString, test.UnEncodedString, encode(test.EncodedString))
			t.Fatal("Not Passing encode() func tests")
		}
	}

}
