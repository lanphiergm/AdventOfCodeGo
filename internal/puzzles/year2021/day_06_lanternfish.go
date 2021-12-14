package year2021

import (
	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Lanternfish Part 1 computes the number of lanternfish after 80 days
func LanternfishPart1(filename string) interface{} {
	return countLanternfish(filename, 80)
}

// Lanternfish Part 2 computes the number of lanternfish after 256 days
func LanternfishPart2(filename string) interface{} {
	return countLanternfish(filename, 256)
}

func countLanternfish(filename string, dayCount int) int {
	fish := utils.ReadCsv(filename)
	var counters [9]int

	// Populate the adult pool from the input
	for _, v := range fish {
		counters[v]++
	}

	// Cycle for the specified number of days
	for i := 0; i < dayCount; i++ {
		generating := counters[0]
		for j := 0; j < 8; j++ {
			counters[j] = counters[j+1]
		}
		counters[6] += generating
		counters[8] = generating
	}

	// Sum up the number of fish
	count := 0
	for i := 0; i < 9; i++ {
		count += counters[i]
	}
	return count
}
