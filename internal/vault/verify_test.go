package vault

import (
	"testing"
)

func TestVerify(t *testing.T) {

	// Erase vault content after the test
	t.Cleanup(cleanup)

	masterkey := "TheMasterKey!"

	// Verify master key against empty vault
	verified, err := Verify(masterkey)
	if err != nil || !verified {
		t.Fatal("Verify against empty vault must return true")
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

	// Verify master key against non-empty vault
	verified, err = Verify(masterkey)
	if err != nil || !verified {
		t.Fatal("Verify against non-empty vault must return true")
	}

	// Verify incorrect master key against non-empty vault
	verified, err = Verify("AnIncorrectMasterKey")
	if err == nil || verified {
		t.Fatal("Verify incorrect master key against non-empty vault must return false, error")
	}
}
