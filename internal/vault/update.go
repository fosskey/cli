package vault

import "errors"

func Update(masterkey, name, secret string) error {

	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return err
	}

	if _, exists := entries[name]; !exists {
		return errors.New("NotFound")
	}

	entries[name] = secret

	// Write to the vault
	if err := write(masterkey, entries); err != nil {
		return err
	}

	return nil
}
