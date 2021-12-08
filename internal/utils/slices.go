package utils

// Removes the element from the specified index. DOES NOT PRESERVE ORDER
func RemoveStr(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
