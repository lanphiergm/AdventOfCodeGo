package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Smoke Basin Part 1 with the sample data
func TestSmokeBasinPart1Sample(t *testing.T) {
	expected := 15
	actual := SmokeBasinPart1("../../../data/year2021/day_09_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Smoke Basin Part 1 with the real puzzle data
func TestSmokeBasinPart1Puzzle(t *testing.T) {
	expected := 558
	actual := SmokeBasinPart1("../../../data/year2021/day_09_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Smoke Basin Part 2 with the sample data
func TestSmokeBasinPart2Sample(t *testing.T) {
	expected := 1134
	actual := SmokeBasinPart2("../../../data/year2021/day_09_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Smoke Basin Part 2 with the real puzzle data
func TestSmokeBasinPart2Puzzle(t *testing.T) {
	expected := 882942
	actual := SmokeBasinPart2("../../../data/year2021/day_09_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
