package utils

import "strconv"

// Converts a string to an int, discarding the error
func Atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

// Converts a rune to an int, discarding the error
func Rtoi(r rune) int {
	return Atoi(string(r))
}

// Converts a byte to an int, discarding the error
func Btoi(b byte) int {
	return Atoi(string(b))
}
