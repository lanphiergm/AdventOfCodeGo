// The 2021 puzzles: https://adventofcode.com/2021
package year2021

import "math"

// Sonar Sweep Part 1 computes the number of times that a depth measurement
// increases.
func sonarSweepPart1(depths []int) int {
	prev := math.MaxInt32
	incCount := 0
	for _, depth := range depths {
		if depth > prev {
			incCount++
		}
		prev = depth
	}
	return incCount
}

// Sonar Sweep Part 2 computes the number of times that a sliding window of
// three consecutive depths increases.
func sonarSweepPart2(depths []int) int {
	incCount := 0
	for i := 1; i < len(depths)-2; i++ {
		sumA := depths[i-1] + depths[i] + depths[i+1]
		sumB := depths[i] + depths[i+1] + depths[i+2]
		if sumB > sumA {
			incCount++
		}
	}
	return incCount
}
