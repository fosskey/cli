package vault

import (
	"github.com/fosskey/cli/internal/cipher"
)

// Verify the master key, return true if verified
// In case of empty vault, return true
func Verify(masterkey string) (bool, error) {

	// Read vault file
	data, err := readVault()
	if err != nil {
		return false, err
	}

	// Return true when vault is empty
	if len(data) == 0 {
		return true, nil
	}

	// Return false if decryption failed
	if _, err := cipher.Decrypt([]byte(masterkey), data); err != nil {
		return false, nil
	}

	// Return true (decryption passed)
	return true, nil
}
