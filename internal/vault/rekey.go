package vault

import "errors"

// Change the master key
func Rekey(masterkey, newkey string) error {

	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return err
	}

	// Return error when vault is empty
	if len(entries) == 0 {
		return errors.New("VaultEmpty")
	}

	// Write to the vault with new masterkey
	if err := write(newkey, entries); err != nil {
		return err
	}

	return nil
}
