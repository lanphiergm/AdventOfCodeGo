package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Transparent Origami Part 1 with the sample data
func TestTransparentOrigamiPart1Sample(t *testing.T) {
	expected := 17
	actual := TransparentOrigamiPart1("../../../data/year2021/day_13_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Transparent Origami Part 1 with the real puzzle data
func TestTransparentOrigamiPart1Puzzle(t *testing.T) {
	expected := 807
	actual := TransparentOrigamiPart1("../../../data/year2021/day_13_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Transparent Origami Part 2 with the real puzzle data
func TestTransparentOrigamiPart2Puzzle(t *testing.T) {
	expected := 98
	actual := TransparentOrigamiPart2("../../../data/year2021/day_13_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
