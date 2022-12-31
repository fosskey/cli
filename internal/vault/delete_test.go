package vault

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Delete something from empty vault
	if err := Delete(masterkey, "Something"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Delete something from an empty vault must return an error(NotFound)")
	}

	// Prepare three entries to be inserted
	entries := map[string]string{
		"FirstSecretName":  "FirstSecretContent",
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
		"FourthSecretName": "FourthSecretContent",
		"FifthSecretName":  "FifthSecretContent",
	}

	// Insert all three entries into the vault
	for name, secret := range entries {
		if err := Insert(masterkey, name, secret); err != nil {
			t.Fatal(err)
		}
	}

	// Delete a non-existent entry from the non-empty vault
	if err := Delete(masterkey, "UnicornEgg"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Delete of a non-existent entry must return an error(NotFound)")
	}

	// Delete an existing entry with an incorrect master key
	if err := Delete("AnIncorrectMasterKey", "FirstSecretName"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to delete with an incorrect master key must return an error(AuthFailed)")
	}

	// Delete a non-existing entry with an incorrect master key
	if err := Delete("AnIncorrectMasterKey", "UnicornEgg"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to delete with an incorrect master key must return an error(AuthFailed)")
	}

	// Delete first entry
	if err := Delete(masterkey, "FirstSecretName"); err != nil {
		t.Fatal(err)
	}

	// Read entries
	readEntries, err := read(masterkey)
	if err != nil || !reflect.DeepEqual(map[string]string{
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
		"FourthSecretName": "FourthSecretContent",
		"FifthSecretName":  "FifthSecretContent",
	}, readEntries) {
		t.Fatal("Read result must reflect the deleted entry")
	}

	// Delete last entry
	if err := Delete(masterkey, "FifthSecretName"); err != nil {
		t.Fatal(err)
	}

	// Read entries
	readEntries, err = read(masterkey)
	if err != nil || !reflect.DeepEqual(map[string]string{
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
		"FourthSecretName": "FourthSecretContent",
	}, readEntries) {
		t.Fatal("Read result must reflect the deleted entry")
	}

	// Delete middle entry
	if err := Delete(masterkey, "ThirdSecretName"); err != nil {
		t.Fatal(err)
	}

	// Read entries
	readEntries, err = read(masterkey)
	if err != nil || !reflect.DeepEqual(map[string]string{
		"SecondSecretName": "SecondSecretContent",
		"FourthSecretName": "FourthSecretContent",
	}, readEntries) {
		t.Fatal("Read result must reflect the deleted entry")
	}
}
