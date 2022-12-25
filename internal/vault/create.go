package vault

import (
	"os"
)

func Create() error {

	fossPath := fossPath()
	vaultPath := vaultPath()

	// Create .foss directory if doesn't exist
	if _, err := os.Stat(fossPath); err != nil {
		if err := os.Mkdir(fossPath, os.ModeDir|os.FileMode(0700)); err != nil {
			return err
		}
	}

	// Create vault file if doesn't exist
	if _, err := os.Stat(vaultPath); err != nil {
		if err := os.WriteFile(vaultPath, nil, os.FileMode(0600)); err != nil {
			return err
		}
	}

	return nil
}
