package vault

import (
	"reflect"
	"testing"
)

func TestStore(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Prepare three entries to be inserted
	entries := map[string]string{
		"FirstSecretName":  "FirstSecretContent",
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
	}

	// Insert all three entries into the vault
	for name, secret := range entries {
		if err := Store(masterkey, name, secret); err != nil {
			t.Fatal(err)
		}
	}

	// Read entries and compare the result
	readEntries, err := read(masterkey)
	if err != nil || !reflect.DeepEqual(entries, readEntries) {
		t.Fatal("Read result must be equal to the inserted entries")
	}

	// Store an entry with an existing name
	if err := Store(masterkey, "FirstSecretName", "Whatever"); err == nil || err.Error() != "DuplicateEntry" {
		t.Fatal("Attempt to store with an existing name must return an error(DuplicateEntry)")
	}

	// Store an entry using an incorrect master key
	if err := Store("AnIncorrectMasterKey", "Something", "Whatever"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to store with an incorrect master key must return an error(AuthFailed)")
	}
}
