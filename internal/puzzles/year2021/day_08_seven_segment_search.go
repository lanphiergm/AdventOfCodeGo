package year2021

import (
	"strconv"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Seven Segment Search Part 1 computes the number of 1s, 4s, 7s, and 8s that
// appear in the output
func SevenSegmentSearchPart1(filename string) int {
	notes := utils.ReadStrings(filename)
	counter := 0
	for _, entry := range notes {
		_, output := parseEntry(entry)
		for _, val := range output {
			if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
				counter++
			}
		}
	}
	return counter
}

// Seven Segment Search Part 2 computes the sum of all output values
func SevenSegmentSearchPart2(filename string) int {
	notes := utils.ReadStrings(filename)
	sum := 0
	for _, entry := range notes {
		sum += decodeEntry(entry)
	}
	return sum
}

func parseEntry(entry string) ([]string, []string) {
	entrySplit := strings.Split(entry, " | ")
	input := strings.Split(entrySplit[0], " ")
	output := strings.Split(entrySplit[1], " ")
	return input, output
}

func decodeEntry(entry string) int {
	var digits [10]string
	var len5 []string
	var len6 []string
	input, output := parseEntry(entry)
	for _, pattern := range input {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 5:
			len5 = append(len5, pattern)
		case 6:
			len6 = append(len6, pattern)
		case 7:
			digits[8] = pattern
		}
	}

	// Find 3 as the only len5 that contains 1
	for i, pattern := range len5 {
		if isSubset(pattern, digits[1]) {
			digits[3] = pattern
			len5 = removeStr(len5, i)
			break
		}
	}

	// Find 9 as the only len6 that contains 3
	for i, pattern := range len6 {
		if isSubset(pattern, digits[3]) {
			digits[9] = pattern
			len6 = removeStr(len6, i)
			break
		}
	}

	// Find 0 as the only remaining len6 that contains 1
	for i, pattern := range len6 {
		if isSubset(pattern, digits[1]) {
			digits[0] = pattern
			len6 = removeStr(len6, i)
			break
		}
	}

	// 6 is the only len6 left
	digits[6] = len6[0]

	// Find 5 as the only len5 that is a subset of 9
	for i, pattern := range len5 {
		if isSubset(digits[9], pattern) {
			digits[5] = pattern
			len5 = removeStr(len5, i)
			break
		}
	}

	// 2 is the only len5 left
	digits[2] = len5[0]

	outputStr := ""
	for _, pattern := range output {
		outputStr += findDigit(pattern, digits)
	}
	return utils.Atoi(outputStr)
}

func isSubset(set string, subset string) bool {
	for _, r := range subset {
		if !strings.ContainsRune(set, r) {
			return false
		}
	}
	return true
}

func areDigitsEqual(a string, b string) bool {
	return len(a) == len(b) && isSubset(a, b)
}

func findDigit(pattern string, digits [10]string) string {
	for i, digit := range digits {
		if areDigitsEqual(pattern, digit) {
			return strconv.Itoa(i)
		}
	}
	return ""
}

func removeStr(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
