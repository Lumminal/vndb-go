package helper

// Contains
//
// Helper to check if a string is contained in a slice/array
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
