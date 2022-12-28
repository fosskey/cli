package vault

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {

	// There is no need to call init() separately
	// Because it will be automatically called before running TestInit

	// Check if foss directory exists
	dir, err := os.Stat(fossPath)
	if err != nil {
		t.Fatal(err)
	}

	// Check dir mode
	dirmode := dir.Mode()
	expected := os.ModeDir | os.FileMode(0700)
	if dirmode != expected {
		t.Fatalf("Expected dir mode %q but found %q", expected, dirmode)
	}

	// Check if vault file exists
	file, err := os.Stat(vaultPath)
	if err != nil {
		t.Fatal(err)
	}

	// Check vault file mode
	filemode := file.Mode()
	expected = os.FileMode(0600)
	if filemode != expected {
		t.Fatalf("Expected file mode %q but found %q", expected, filemode)
	}
}
