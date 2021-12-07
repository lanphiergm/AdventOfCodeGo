package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests The Treachery of Whales Part 1 with the sample data
func TestTheTreacheryofWhalesPart1Sample(t *testing.T) {
	expected := 37
	actual := TheTreacheryofWhalesPart1("../../../data/year2021/day_07_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests The Treachery of Whales Part 1 with the real puzzle data
func TestTheTreacheryofWhalesPart1Puzzle(t *testing.T) {
	expected := 342641
	actual := TheTreacheryofWhalesPart1("../../../data/year2021/day_07_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests The Treachery of Whales Part 2 with the sample data
func TestTheTreacheryofWhalesPart2Sample(t *testing.T) {
	expected := 168
	actual := TheTreacheryofWhalesPart2("../../../data/year2021/day_07_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests The Treachery of Whales Part 2 with the real puzzle data
func TestTheTreacheryofWhalesPart2Puzzle(t *testing.T) {
	expected := 93006301
	actual := TheTreacheryofWhalesPart2("../../../data/year2021/day_07_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
