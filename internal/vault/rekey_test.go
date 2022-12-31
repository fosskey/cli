package vault

import (
	"reflect"
	"testing"
)

func TestRekey(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Rekey on empty vault
	if err := Rekey(masterkey, "Whatever"); err == nil || err.Error() != "VaultEmpty" {
		t.Fatal("Rekey on empty vault must return an error(VaultEmpty)")
	}

	// Prepare three entries to be inserted
	entries := map[string]string{
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

	// Rekey on non-empty vault with an incorrect master key
	if err := Rekey("AnIncorrectMasterKey", "Whatever"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Rekey on non-empty vault with an incorrect master key must return an error(AuthFailed)")
	}

	// Rekey on non-empty vault with the correct master key
	newkey := "NewMasterKey!"
	if err := Rekey(masterkey, newkey); err != nil {
		t.Fatal(err)
	}

	// Read with old master key
	if _, err := read(masterkey); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to read with old master key must return an error(AuthFailed)")
	}

	// Read with new masterkey and compare the result
	readEntries, err := read(newkey)
	if err != nil || !reflect.DeepEqual(entries, readEntries) {
		t.Fatal("Read result must be equal to the inserted entries")
	}
}
