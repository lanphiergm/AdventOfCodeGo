package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests HydrothermalVenture Part 1 with the sample data
func TestHydrothermalVenturePart1Sample(t *testing.T) {
	expected := 5
	actual := HydrothermalVenturePart1("../../../data/year2021/day_05_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests HydrothermalVenture Part 1 with the real puzzle data
func TestHydrothermalVenturePart1Puzzle(t *testing.T) {
	expected := 6113
	actual := HydrothermalVenturePart1("../../../data/year2021/day_05_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests HydrothermalVenture Part 2 with the sample data
func TestHydrothermalVenturePart2Sample(t *testing.T) {
	expected := 12
	actual := HydrothermalVenturePart2("../../../data/year2021/day_05_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests HydrothermalVenture Part 2 with the real puzzle data
func TestHydrothermalVenturePart2Puzzle(t *testing.T) {
	expected := 0
	actual := HydrothermalVenturePart2("../../../data/year2021/day_05_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
