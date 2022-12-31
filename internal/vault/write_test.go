package vault

import (
	"os"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Write empty entries to the vault
	entries := map[string]string{}
	if err := write(masterkey, entries); err != nil {
		t.Fatal(err)
	}

	// Read entries and compare the result
	readEntries, err := read(masterkey)
	if err != nil || len(readEntries) != 0 {
		t.Fatal("Read result must be an empty map")
	}

	// Prepare three entries to be inserted
	entries = map[string]string{
		"FirstSecretName":  "FirstSecretContent",
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
	}

	// Write to the vault
	if err := write(masterkey, entries); err != nil {
		t.Fatal(err)
	}

	// Get vault file info
	file, err := os.Stat(vaultPath)
	if err != nil {
		t.Fatal(err)
	}

	// Check vault file size
	if file.Size() == 0 {
		t.Fatal("Vault file is empty after write")
	}

	// Check vault file mode
	filemode := file.Mode()
	expected := os.FileMode(0600)
	if filemode != expected {
		t.Fatalf("Expected file mode %q but found %q", expected, filemode)
	}

	// Read entries and compare the result
	readEntries, err = read(masterkey)
	if err != nil || !reflect.DeepEqual(entries, readEntries) {
		t.Fatal("Read result must be equal to the inserted entries")
	}
}
