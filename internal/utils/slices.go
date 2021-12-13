package utils

// Returns only the distinct values from the specified slice
func DistinctStr(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Removes the element from the specified index. DOES NOT PRESERVE ORDER
func RemoveStr(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
