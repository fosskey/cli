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
