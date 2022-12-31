package vault

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Read entries from empty vault
	entries, err := read(masterkey)
	if err != nil || len(entries) != 0 {
		t.Fatal("Read from empty vault must return an empty map")
	}

	// Prepare three entries to be inserted
	entries = map[string]string{
		"FirstSecretName":  "FirstSecretContent",
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
	}

	// Insert all three entries into the vault
	for name, secret := range entries {
		if err := Insert(masterkey, name, secret); err != nil {
			t.Fatal(err)
		}
	}

	// Read and compare the result
	readEntries, err := read(masterkey)
	if err != nil || !reflect.DeepEqual(entries, readEntries) {
		t.Fatal("Read result must be equal to the inserted entries")
	}

	// Read with an incorrect master key
	if _, err := read("AnIncorrectMasterKey"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to read with an incorrect master key must return an error(AuthFailed)")
	}
}
