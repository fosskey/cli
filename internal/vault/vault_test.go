package vault

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestStoreFetch(t *testing.T) {

	fossDir = ".fosstest"
	fossPath := fossPath()
	vaultPath := vaultPath()

	// Cleanup will be called after all tests complete
	t.Cleanup(func() {
		if err := os.RemoveAll(fossPath); err != nil {
			t.Fatal(err)
		}
	})

	masterkey := "TheMasterKey!"

	// Verify master key before creating the vault
	verified, err := Verify(masterkey)
	if err != nil || !verified {
		t.Fatal("Verify before creating vault must return true")
	}

	// Fetch all before creating the vault
	entries, err := FetchAll(masterkey)
	if err != nil || len(entries) != 0 {
		t.Fatal("Fetch all before creating vault must return an empty map")
	}

	// Fetch something before creating the vault
	if _, err := Fetch(masterkey, "Something"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Fetch something before creating vault must return an error(NotFound)")
	}

	// List before creating the vault
	names, err := List(masterkey)
	if err != nil || len(names) != 0 {
		t.Fatal("List before creating vault must return 0")
	}

	// Run Create func first
	if err := Create(); err != nil {
		t.Fatal(err)
	}

	// Verify master key against empty vault
	verified, err = Verify(masterkey)
	if err != nil || !verified {
		t.Fatal("Verify against empty vault must return true")
	}

	// Fetch all from empty vault
	entries, err = FetchAll(masterkey)
	if err != nil || len(entries) != 0 {
		t.Fatal("Fetch all from empty vault must return an empty map")
	}

	// Fetch something from empty vault
	if _, err := Fetch(masterkey, "Something"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Fetch something from empty vault must return an error(NotFound)")
	}

	// List from empty vault
	names, err = List(masterkey)
	if err != nil || len(names) != 0 {
		t.Fatal("List from empty vault must return 0")
	}

	// Prepare three entries to be inserted
	entries = make(map[string]string)
	entries["FirstSecretName"] = "FirstSecretContent"
	entries["SecondSecretName"] = "SecondSecretContent"
	entries["ThirdSecretName"] = "ThirdSecretContent"

	// Get the entry keys (names) for future use
	entryNames := []string{}
	for k := range entries {
		entryNames = append(entryNames, k)
	}
	sort.Strings(entryNames) // sort

	// Insert all three entries to the vault
	for name, secret := range entries {
		if err := Store(masterkey, name, secret); err != nil {
			t.Fatal(err)
		}
	}

	// Get file info
	file, err := os.Stat(vaultPath)
	if err != nil {
		t.Fatal(err)
	}

	// Check file size
	if file.Size() == 0 {
		t.Fatal("Vault file is empty after store")
	}

	// Check file mode
	filemode := file.Mode()
	expected := os.FileMode(0600)
	if filemode != expected {
		t.Fatalf("Expected file mode %q but found %q", expected, filemode)
	}

	// Verify master key against non-empty vault
	verified, err = Verify(masterkey)
	if err != nil || !verified {
		t.Fatal("Verify against non-empty vault must return true")
	}

	// Fetch a non-existent entry from a non-empty vault
	if _, err := Fetch(masterkey, "Something"); err == nil || err.Error() != "NotFound" {
		t.Fatal("Fetch result for a non-existent entry must return an error(NotFound)")
	}

	// Fetch and check all three entries
	for name, secret := range entries {
		fetchedSecret, err := Fetch(masterkey, name)
		if err != nil {
			t.Fatal(err)
		}
		if fetchedSecret != secret {
			t.Fatalf("Expected %q, got %q", secret, fetchedSecret)
		}
	}

	// List from non-empty vault
	names, err = List(masterkey)
	if err != nil || !reflect.DeepEqual(names, entryNames) {
		t.Fatal("List from non-empty vault must return all names")
	}

	// Store an entry with existing name
	if err := Store(masterkey, "FirstSecretName", "Whatever"); err == nil || err.Error() != "DuplicateEntry" {
		t.Fatal("Attempt to store with an existing name must return an error(DuplicateEntry)")
	}

	// Verify incorrect master key against non-empty vault
	verified, err = Verify("AnIncorrectMasterKey")
	if err != nil || verified {
		t.Fatal("Verify incorrect master key against non-empty vault must return false")
	}

	// Fetch an existing entry with an incorrect master key
	if _, err := Fetch("AnIncorrectMasterKey", "FirstSecretName"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to fetch with an incorrect master key must return an error(AuthFailed)")
	}

	// Fetch a non-existing entry with an incorrect master key
	if _, err := Fetch("AnIncorrectMasterKey", "Something"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to fetch with an incorrect master key must return an error(AuthFailed)")
	}

	// List with an incorrect master key
	if _, err := List("AnIncorrectMasterKey"); err == nil || err.Error() != "AuthFailed" {
		t.Fatal("Attempt to list with an incorrect master key must return an error(AuthFailed)")
	}
}
