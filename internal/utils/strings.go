package utils

import "strconv"

// Converts a string to an int, discarding the error
func Atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
