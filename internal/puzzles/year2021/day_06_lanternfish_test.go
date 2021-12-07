package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Lanternfish Part 1 with the sample data
func TestLanternfishPart1Sample(t *testing.T) {
	expected := 5934
	actual := LanternfishPart1("../../../data/year2021/day_06_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Lanternfish Part 1 with the real puzzle data
func TestLanternfishPart1Puzzle(t *testing.T) {
	expected := 386640
	actual := LanternfishPart1("../../../data/year2021/day_06_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Lanternfish Part 2 with the sample data
func TestLanternfishPart2Sample(t *testing.T) {
	expected := 26984457539
	actual := LanternfishPart2("../../../data/year2021/day_06_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Lanternfish Part 2 with the real puzzle data
func TestLanternfishPart2Puzzle(t *testing.T) {
	expected := 1733403626279
	actual := LanternfishPart2("../../../data/year2021/day_06_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
