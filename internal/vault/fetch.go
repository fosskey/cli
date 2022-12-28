package vault

import (
	"errors"
)

func Fetch(masterkey, name string) (string, error) {

	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return "", err
	}

	if _, exists := entries[name]; exists {
		return entries[name], nil
	}

	return "", errors.New("NotFound")
}
