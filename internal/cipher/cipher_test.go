package cipher

import (
	"bytes"
	"testing"
)

func TestCipher(t *testing.T) {
	var key = []byte("")
	var message = []byte("Hello World!")

	ciphertext, err := Encrypt(key, message)
	if err != nil {
		t.Fatal(err)
	}

	if len(ciphertext) <= len(message) {
		t.Fatal("Encrypted result must be longer than the original message")
	}

	plaintext, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(plaintext, message) {
		t.Fatal("Decrypted result did not match with original message")
	}
}
