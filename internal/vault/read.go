package vault

import (
	"errors"
	"os"
	"strings"

	"github.com/fosskey/cli/internal/cipher"
)

// Read and decrypt the content of the vault
func read(masterkey string) (map[string]string, error) {

	// Read vault file
	encryptedBytes, err := os.ReadFile(vaultPath)
	if err != nil {
		return nil, err
	}

	// Return empty map when vault is empty
	if len(encryptedBytes) == 0 {
		return make(map[string]string), nil
	}

	// Decrypt
	decryptedBytes, err := cipher.Decrypt([]byte(masterkey), encryptedBytes)
	if err != nil {
		return nil, errors.New("AuthFailed")
	}

	// Convert to string
	plainText := string(decryptedBytes)

	// Map the plain text into key:value entries
	entries := make(map[string]string)
	lines := strings.Split(plainText, "\n")
	for _, line := range lines {
		v := strings.Split(line, "\t")
		entries[v[0]] = v[1]
	}

	return entries, nil
}
