package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Snailfish Part 1 with the sample data
func TestSnailfishPart1Sample(t *testing.T) {
	expected := 4140
	actual := SnailfishPart1("../../../data/year2021/day_18_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Snailfish Part 1 with the real puzzle data
func TestSnailfishPart1Puzzle(t *testing.T) {
	expected := 3734
	actual := SnailfishPart1("../../../data/year2021/day_18_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Snailfish Part 2 with the sample data
func TestSnailfishPart2Sample(t *testing.T) {
	expected := 3993
	actual := SnailfishPart2("../../../data/year2021/day_18_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Snailfish Part 2 with the real puzzle data
func TestSnailfishPart2Puzzle(t *testing.T) {
	expected := 4837
	actual := SnailfishPart2("../../../data/year2021/day_18_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
