package vault

import (
	"errors"
	"os"
	"strings"

	"github.com/fosskey/cli/internal/cipher"
)

func Store(masterkey, name, secret string) error {

	vaultPath := vaultPath()

	// Fetch all
	entries, err := FetchAll(masterkey)
	if err != nil {
		return err
	}

	// Check if the name already exists
	if _, exists := entries[name]; exists {
		return errors.New("DuplicateEntry")
	}

	// Append the new entry
	entries[name] = secret

	// Flatten all entries into plain text
	content := ""
	for name, secret := range entries {
		content += name + "\t" + secret + "\n"
	}

	// Trim the final "\n"
	content = strings.Trim(content, "\n")

	// Encrypt
	encryptedData, err := cipher.Encrypt([]byte(masterkey), []byte(content))
	if err != nil {
		return err
	}

	// Write to vault
	if err := os.WriteFile(vaultPath, encryptedData, 0); err != nil {
		return err
	}

	return nil
}
