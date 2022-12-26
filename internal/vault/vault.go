package vault

import (
	"os"
	"path/filepath"
)

var fossDir = ".foss"

func fossPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, fossDir)
}

func vaultPath() string {
	return filepath.Join(fossPath(), "vault")
}

func readVault() ([]byte, error) {
	vaultPath := vaultPath()

	// Create vault if doesn't exist
	if err := Create(); err != nil {
		return nil, err
	}

	// Read vault file
	data, err := os.ReadFile(vaultPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
