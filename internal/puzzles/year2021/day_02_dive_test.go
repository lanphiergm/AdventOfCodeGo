package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Dive Part 1 with the sample data
func TestDivePart1Sample(t *testing.T) {
	expected := 150
	actual := DivePart1("../../../data/year2021/day_02_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dive Part 1 with the puzzle data
func TestDivePart1Puzzle(t *testing.T) {
	expected := 1484118
	actual := DivePart1("../../../data/year2021/day_02_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dive Part 2 with the sample data
func TestDivePart2Sample(t *testing.T) {
	expected := 900
	actual := DivePart2("../../../data/year2021/day_02_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dive Part 2 with the puzzle data
func TestDivePart2Puzzle(t *testing.T) {
	expected := 1463827010
	actual := DivePart2("../../../data/year2021/day_02_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
