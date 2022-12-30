package vault

import (
	"testing"
)

func TestUpdate(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Update something from empty vault
	if err := Update(masterkey, "Something", "Whatever"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Update something in an empty vault must return an error(NotFound)")
	}

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

	// Update a non-existent entry in the non-empty vault
	if err := Update(masterkey, "UnicornEgg", "Whatever"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Update of a non-existent entry must return an error(NotFound)")
	}

	// Update all three entries by appending "Updated" to the end
	for name, secret := range entries {
		err := Update(masterkey, name, secret+"Updated")
		if err != nil {
			t.Fatal(err)
		}
	}

	// Fetch and compare all three entries
	for name, secret := range entries {
		fetchedSecret, err := Fetch(masterkey, name)
		if err != nil {
			t.Fatal(err)
		}
		if fetchedSecret != secret+"Updated" {
			t.Fatalf("Expected %qUpdated, got %q", secret, fetchedSecret)
		}
	}
}
