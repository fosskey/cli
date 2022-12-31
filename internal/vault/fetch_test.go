package vault

import (
	"testing"
)

func TestFetch(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Fetch something from empty vault
	if _, err := Fetch(masterkey, "Something"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Fetch something from empty vault must return an error(NotFound)")
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

	// Fetch and compare all three entries
	for name, secret := range entries {
		fetchedSecret, err := Fetch(masterkey, name)
		if err != nil {
			t.Fatal(err)
		}
		if fetchedSecret != secret {
			t.Fatalf("Expected %q, got %q", secret, fetchedSecret)
		}
	}

	// Fetch a non-existent entry from the vault
	if _, err := Fetch(masterkey, "UnicornEgg"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Fetch of a non-existent entry must return an error(NotFound)")
	}

	// Fetch an existing entry with an incorrect master key
	if _, err := Fetch("AnIncorrectMasterKey", "FirstSecretName"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to fetch with an incorrect master key must return an error(AuthFailed)")
	}

	// Fetch a non-existing entry with an incorrect master key
	if _, err := Fetch("AnIncorrectMasterKey", "UnicornEgg"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to fetch with an incorrect master key must return an error(AuthFailed)")
	}
}
