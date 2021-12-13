package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Passage Pathing Part 1 with the sample data
func TestPassagePathingPart1Sample(t *testing.T) {
	expected := 226
	actual := PassagePathingPart1("../../../data/year2021/day_12_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Passage Pathing Part 1 with the real puzzle data
func TestPassagePathingPart1Puzzle(t *testing.T) {
	expected := 5178
	actual := PassagePathingPart1("../../../data/year2021/day_12_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Passage Pathing Part 2 with the sample data
func TestPassagePathingPart2Sample(t *testing.T) {
	expected := 3509
	actual := PassagePathingPart2("../../../data/year2021/day_12_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Passage Pathing Part 2 with the real puzzle data
func TestPassagePathingPart2Puzzle(t *testing.T) {
	expected := 130094
	actual := PassagePathingPart2("../../../data/year2021/day_12_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
