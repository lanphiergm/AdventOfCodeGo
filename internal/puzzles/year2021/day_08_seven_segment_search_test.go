package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Seven Segment Search Part 1 with the sample data
func TestSevenSegmentSearchPart1Sample(t *testing.T) {
	expected := 26
	actual := SevenSegmentSearchPart1("../../../data/year2021/day_08_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Seven Segment Search Part 1 with the real puzzle data
func TestSevenSegmentSearchPart1Puzzle(t *testing.T) {
	expected := 344
	actual := SevenSegmentSearchPart1("../../../data/year2021/day_08_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Seven Segment Search Part 2 with the sample data
func TestSevenSegmentSearchPart2Sample(t *testing.T) {
	expected := 61229
	actual := SevenSegmentSearchPart2("../../../data/year2021/day_08_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Seven Segment Search Part 2 with the real puzzle data
func TestSevenSegmentSearchPart2Puzzle(t *testing.T) {
	expected := 1048410
	actual := SevenSegmentSearchPart2("../../../data/year2021/day_08_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
