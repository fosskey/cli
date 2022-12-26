package vault

import "sort"

// List all names
func List(masterkey string) ([]string, error) {

	// Fetch all entries
	entries, err := FetchAll(masterkey)
	if err != nil {
		return nil, err
	}

	names := []string{}

	// Iterate thru entries and get the names
	for k := range entries {
		names = append(names, k)
	}

	// Sort the slice
	sort.Strings(names)

	return names, nil
}
