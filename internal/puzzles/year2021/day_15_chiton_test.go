package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Chiton Part 1 with the sample data
func TestChitonPart1Sample(t *testing.T) {
	expected := 40
	actual := ChitonPart1("../../../data/year2021/day_15_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Chiton Part 1 with the real puzzle data
func TestChitonPart1Puzzle(t *testing.T) {
	expected := 720
	actual := ChitonPart1("../../../data/year2021/day_15_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Chiton Part 2 with the sample data
func TestChitonPart2Sample(t *testing.T) {
	expected := 315
	actual := ChitonPart2("../../../data/year2021/day_15_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Chiton Part 2 with the real puzzle data
func TestChitonPart2Puzzle(t *testing.T) {
	expected := 3025
	actual := ChitonPart2("../../../data/year2021/day_15_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
