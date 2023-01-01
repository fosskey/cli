package cipher

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

// Argon2id password hashing config.
// Based on the recommended parameters from RFC 9106:
// https://www.rfc-editor.org/rfc/rfc9106.html#name-parameter-choice
var config = struct {
	time    uint32 // number of passes
	memory  uint32 // memory size in KiB
	threads uint8  // degree of parallelism
	taglen  uint32 // tag length in bytes
	saltlen int    // salt length in bytes
}{
	time:    1,               // 1 pass
	memory:  2 * 1024 * 1024, // 2 GiB
	threads: 4,               // 4 lanes
	taglen:  32,              // 256-bit tag
	saltlen: 16,              // 128-bit salt
}

func Encrypt(password, plaintext []byte) ([]byte, error) {

	// Generate random salt and Argon2id key from the password
	key, salt, err := deriveKey(password, nil)
	if err != nil {
		return nil, err
	}

	// Get XChaCha20-Poly1305 AEAD from the key
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	// Select a random nonce, and leave capacity for the ciphertext
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(plaintext)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	// Encrypt the message and append the ciphertext to the nonce
	ciphertext := aead.Seal(nonce, nonce, plaintext, nil)

	// Append the salt
	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(password, ciphertext []byte) ([]byte, error) {

	// Split salt and ciphertext
	salt, ciphertext := ciphertext[len(ciphertext)-config.saltlen:], ciphertext[:len(ciphertext)-config.saltlen]

	// Derive key from salt
	key, _, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}

	// Get XChaCha20-Poly1305 AEAD from the key
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aead.NonceSize() {
		return nil, errors.New("cipher text too short")
	}

	// Split nonce and ciphertext
	nonce, ciphertext := ciphertext[:aead.NonceSize()], ciphertext[aead.NonceSize():]

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
		salt = make([]byte, config.saltlen)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key := argon2.IDKey(password, salt, config.time, config.memory, config.threads, config.taglen)

	return key, salt, nil
}
