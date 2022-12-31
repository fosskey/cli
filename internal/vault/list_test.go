package vault

import (
	"reflect"
	"sort"
	"testing"
)

func TestList(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// List from empty vault
	names, err := List(masterkey)
	if err != nil || len(names) != 0 {
		t.Fatal("List from empty vault must return an empty slice")
	}

	// Prepare three entries to be inserted
	entries := map[string]string{
		"FirstSecretName":  "FirstSecretContent",
		"SecondSecretName": "SecondSecretContent",
		"ThirdSecretName":  "ThirdSecretContent",
	}

	// Get the entry names (keys) for future use
	entryNames := []string{}
	for k := range entries {
		entryNames = append(entryNames, k)
	}
	sort.Strings(entryNames) // Sort for future comparison

	// Insert all three entries into the vault
	for name, secret := range entries {
		if err := Insert(masterkey, name, secret); err != nil {
			t.Fatal(err)
		}
	}

	// List from non-empty vault
	names, err = List(masterkey)
	if err != nil || !reflect.DeepEqual(names, entryNames) {
		t.Fatal("List from non-empty vault must return all names")
	}

	// List with an incorrect master key
	if _, err := List("AnIncorrectMasterKey"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to list with an incorrect master key must return an error(AuthFailed)")
	}
}
