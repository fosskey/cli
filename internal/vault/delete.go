package vault

import "errors"

func Delete(masterkey, name string) error {
	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return err
	}

	if _, exists := entries[name]; !exists {
		return errors.New("NotFound")
	}

	delete(entries, name)

	// Write to the vault
	if err := write(masterkey, entries); err != nil {
		return err
	}

	return nil
}
