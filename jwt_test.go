package jwt

import (
	"testing"
)

func TestPrivatehash(t *testing.T) {

	if !IsPrivateKeySet() {
		SetPrivateKey("testingprivatekey")
	}

	h, _ := hashTheString("Muthu")
	t.Logf("Hash: %s\n", h)

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
