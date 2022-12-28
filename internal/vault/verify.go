package vault

// Verify the master key, return true if verified
// In case of empty vault, return true
func Verify(masterkey string) (bool, error) {

	// Read entries
	entries, err := read(masterkey)
	if err != nil {
		return false, err
	}

	// Return true when vault is empty
	if len(entries) == 0 {
		return true, nil
	}

	// Return true (decryption passed)
	return true, nil
}
