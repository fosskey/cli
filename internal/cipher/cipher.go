package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

func Encrypt(password, data []byte) ([]byte, error) {
	key, salt, err := deriveKey(password, nil)
	if err != nil {
		return nil, err
	}

	// Generate a new AES-256 cipher block using our 32-byte key
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	// Creates a new byte array of the size as nonce which will be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// Populates our nonce with a cryptographically secure random sequence
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(password, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, _, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}

	// Generate a new AES-256 cipher block using our 32-byte key
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func deriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, 32768, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}
