package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Trick Shot Part 1 with the sample data
func TestTrickShotPart1Sample(t *testing.T) {
	expected := 45
	actual := TrickShotPart1("../../../data/year2021/day_17_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Trick Shot Part 1 with the real puzzle data
func TestTrickShotPart1Puzzle(t *testing.T) {
	expected := 5886
	actual := TrickShotPart1("../../../data/year2021/day_17_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Trick Shot Part 2 with the sample data
func TestTrickShotPart2Sample(t *testing.T) {
	expected := 112
	actual := TrickShotPart2("../../../data/year2021/day_17_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Trick Shot Part 2 with the real puzzle data
func TestTrickShotPart2Puzzle(t *testing.T) {
	expected := 1806
	actual := TrickShotPart2("../../../data/year2021/day_17_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
