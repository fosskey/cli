package vault

import (
	"errors"
)

func Insert(masterkey, name, secret string) error {

	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return err
	}

	// Check if the name already exists
	if _, exists := entries[name]; exists {
		return errors.New("DuplicateEntry")
	}

	// Append the new entry
	entries[name] = secret

	// Write to the vault
	if err := write(masterkey, entries); err != nil {
		return err
	}

	return nil
}
