package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Extended Polymerization Part 1 with the sample data
func TestExtendedPolymerizationPart1Sample(t *testing.T) {
	expected := 1588
	actual := ExtendedPolymerizationPart1("../../../data/year2021/day_14_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Extended Polymerization Part 1 with the real puzzle data
func TestExtendedPolymerizationPart1Puzzle(t *testing.T) {
	expected := 2915
	actual := ExtendedPolymerizationPart1("../../../data/year2021/day_14_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Extended Polymerization Part 2 with the sample data
func TestExtendedPolymerizationPart2Sample(t *testing.T) {
	expected := 2188189693529
	actual := ExtendedPolymerizationPart2("../../../data/year2021/day_14_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Extended Polymerization Part 2 with the real puzzle data
func TestExtendedPolymerizationPart2Puzzle(t *testing.T) {
	expected := 3353146900153
	actual := ExtendedPolymerizationPart2("../../../data/year2021/day_14_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
