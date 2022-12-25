package vault

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {

	fossDir = ".fosstest"
	fossPath := fossPath()
	vaultPath := vaultPath()

	// Cleanup will be called after all tests complete
	t.Cleanup(func() {
		if err := os.RemoveAll(fossPath); err != nil {
			t.Fatal(err)
		}
	})

	// Run Create func
	if err := Create(); err != nil {
		t.Fatal(err)
	}

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

	// Check file mode
	filemode := file.Mode()
	expected = os.FileMode(0600)
	if filemode != expected {
		t.Fatalf("Expected file mode %q but found %q", expected, filemode)
	}

}
