package vault

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	fossPath  string
	vaultPath string
)

// Create the vault file inside foss directory if it doesn't exist.
func init() {

	// Set paths
	setPaths()

	// Create foss directory if doesn't exist
	if _, err := os.Stat(fossPath); err != nil {
		err := os.Mkdir(fossPath, os.ModeDir|os.FileMode(0700))
		handleErr(err)
	}

	// If the vault file doesn't exist, create it.
	if _, err := os.Stat(vaultPath); err != nil {
		err := os.WriteFile(vaultPath, nil, os.FileMode(0600))
		handleErr(err)
	}
}

// For tests set fossPath to "<ProjectRoot>/.foss", otherwise "<HomeDir>/.foss"
// And set vaultPath to "<fossPath>/vault"
func setPaths() {
	if strings.HasSuffix(os.Args[0], ".test") {
		// Test case
		f, _ := os.Getwd()
		rootPath := path.Join(f, "..", "..")
		fossPath = filepath.Join(rootPath, ".foss")
	} else {
		// Run case
		homePath, _ := os.UserHomeDir()
		fossPath = filepath.Join(homePath, ".foss")
	}
	vaultPath = filepath.Join(fossPath, "vault")
}

// Clean up vault file
func cleanup() {
	err := os.WriteFile(vaultPath, nil, 0)
	handleErr(err)
}

// Print the error with the prefix "Error:" and exit with error code 1.
// It does nothing if the error is nil.
func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
