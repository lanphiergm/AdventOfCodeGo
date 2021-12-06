package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests GiantSquid Part 1 with the sample data
func TestGiantSquidPart1Sample(t *testing.T) {
	expected := 4512
	actual := GiantSquidPart1("../../../data/year2021/day_04_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests GiantSquid Part 1 with the real puzzle data
func TestGiantSquidPart1Puzzle(t *testing.T) {
	expected := 11774
	actual := GiantSquidPart1("../../../data/year2021/day_04_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests GiantSquid Part 2 with the sample data
func TestGiantSquidPart2Sample(t *testing.T) {
	expected := 1924
	actual := GiantSquidPart2("../../../data/year2021/day_04_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests GiantSquid Part 2 with the real puzzle data
func TestGiantSquidPart2Puzzle(t *testing.T) {
	expected := 4495
	actual := GiantSquidPart2("../../../data/year2021/day_04_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
