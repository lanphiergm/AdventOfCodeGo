package year2021

import (
	"math"
	"strconv"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Binary Diagnostic Part 1 computes the power consumption of the submarine
func BinaryDiagnosticPart1(filename string) interface{} {
	report := utils.ReadStrings(filename)
	bits := len(report[0])
	bitCounts := make([]int, bits)

	for _, line := range report {
		runes := []rune(line)
		for i, r := range runes {
			if r == '1' {
				bitCounts[i]++
			}
		}
	}

	reportMajority := math.Ceil(float64(len(report)) / 2.0)
	var gammaStr string
	var epsilonStr string
	for i := 0; i < bits; i++ {
		if bitCounts[i] >= int(reportMajority) {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 32)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 32)
	return int(gamma) * int(epsilon)
}

// Binary Diagnostic Part 2 computes the life support rating of the submarine
func BinaryDiagnosticPart2(filename string) interface{} {
	report := utils.ReadStrings(filename)
	oxygenGeneratorRating := GetRating(report, true, '1')
	co2ScrubberRating := GetRating(report, false, '0')
	return oxygenGeneratorRating * co2ScrubberRating
}

func GetRating(report []string, keepMostCommon bool, tieRune rune) int {
	var ratingStr string
	reportCopy := make([]string, len(report))
	copy(reportCopy, report)
	bits := len(reportCopy[0])
	for i := 0; i < bits; i++ {
		// Count the set bits at the current index
		bitCount := 0
		for _, r := range reportCopy {
			if []rune(r)[i] == '1' {
				bitCount++
			}
		}

		// Figure out which character to keep
		if len(reportCopy)-bitCount == bitCount {
			ratingStr += string(tieRune)
		} else if len(reportCopy)-bitCount < bitCount {
			if keepMostCommon {
				ratingStr += "1"
			} else {
				ratingStr += "0"
			}
		} else {
			if keepMostCommon {
				ratingStr += "0"
			} else {
				ratingStr += "1"
			}
		}

		// Filter out the ones we don't want to keep
		for j := 0; j < len(reportCopy); {
			if strings.HasPrefix(reportCopy[j], ratingStr) {
				j++
			} else {
				reportCopy = utils.RemoveStr(reportCopy, j)
			}
		}

		if len(reportCopy) == 1 {
			ratingStr = reportCopy[0]
			break
		}
	}

	rating, _ := strconv.ParseInt(ratingStr, 2, 32)
	return int(rating)
}
