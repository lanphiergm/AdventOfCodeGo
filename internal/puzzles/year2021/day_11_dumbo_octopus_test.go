package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Dumbo Octopus Part 1 with the sample data
func TestDumboOctopusPart1Sample(t *testing.T) {
	expected := 1656
	actual := DumboOctopusPart1("../../../data/year2021/day_11_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dumbo Octopus Part 1 with the real puzzle data
func TestDumboOctopusPart1Puzzle(t *testing.T) {
	expected := 1601
	actual := DumboOctopusPart1("../../../data/year2021/day_11_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dumbo Octopus Part 2 with the sample data
func TestDumboOctopusPart2Sample(t *testing.T) {
	expected := 195
	actual := DumboOctopusPart2("../../../data/year2021/day_11_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Dumbo Octopus Part 2 with the real puzzle data
func TestDumboOctopusPart2Puzzle(t *testing.T) {
	expected := 368
	actual := DumboOctopusPart2("../../../data/year2021/day_11_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
