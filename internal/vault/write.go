package vault

import (
	"os"
	"strings"

	"github.com/fosskey/cli/internal/cipher"
)

// Encrypt entries and write them into the vault
func write(masterkey string, entries map[string]string) error {

	// Return right away if entries is an empty map
	if len(entries) == 0 {
		return nil
	}

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

	// Write to the vault
	if err := os.WriteFile(vaultPath, encryptedData, 0); err != nil {
		return err
	}

	return nil
}
