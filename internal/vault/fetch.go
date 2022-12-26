package vault

import (
	"errors"
	"strings"

	"github.com/fosskey/cli/internal/cipher"
)

func Fetch(masterkey, name string) (string, error) {

	entries, err := FetchAll(masterkey)
	if err != nil {
		return "", err
	}

	if _, exists := entries[name]; exists {
		return entries[name], nil
	}

	return "", errors.New("NotFound")
}

func FetchAll(masterkey string) (map[string]string, error) {

	// Read vault file
	data, err := readVault()
	if err != nil {
		return nil, err
	}

	// Return empty map when vault is empty
	if len(data) == 0 {
		return make(map[string]string), nil
	}

	// Decrypt
	bytes, err := cipher.Decrypt([]byte(masterkey), data)
	if err != nil {
		return nil, errors.New("AuthFailed")
	}

	// Convert to string
	content := string(bytes)

	// Map the content into key:value entries
	entries := make(map[string]string)
	lines := strings.Split(string(content[:]), "\n")
	for _, line := range lines {
		v := strings.Split(line, "\t")
		entries[v[0]] = v[1]
	}

	return entries, nil
}
