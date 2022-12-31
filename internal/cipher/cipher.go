package cipher

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

// Argon2id password hashing config
var config = struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}{
	time:    1,         // number of passes
	memory:  64 * 1024, // 64 MB
	threads: 4,         // numbers of available CPUs
	keyLen:  32,        // 32 bytes (256-bit)
}

func Encrypt(password, data []byte) ([]byte, error) {

	// Generate Argon2id key and salt using the password
	key, salt, err := deriveKey(password, nil)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	// Select a random nonce, and leave capacity for the ciphertext
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(data)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	// Encrypt the message and append the ciphertext to the nonce
	ciphertext := aead.Seal(nonce, nonce, data, nil)

	// Append the salt
	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(password, ciphertext []byte) ([]byte, error) {

	// Extract salt and data from ciphertext
	salt, data := ciphertext[len(ciphertext)-32:], ciphertext[:len(ciphertext)-32]

	key, _, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aead.NonceSize() {
		return nil, errors.New("cipher text too short")
	}

	// Split nonce and ciphertext
	nonce, ciphertext := data[:aead.NonceSize()], data[aead.NonceSize():]

	// Decrypt the message and check it wasn't tampered with
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Derives a key from the password, salt, and cost parameters using Argon2id
func deriveKey(password, salt []byte) ([]byte, []byte, error) {

	if salt == nil {
		// Generate a salt
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key := argon2.IDKey(password, salt, config.time, config.memory, config.threads, config.keyLen)

	return key, salt, nil
}
