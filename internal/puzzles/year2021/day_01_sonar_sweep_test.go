package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Sonar Sweep Part 1 with the sample data
func TestSonarSweepPart1Sample(t *testing.T) {
	expected := 7
	actual := SonarSweepPart1("../../../data/year2021/day_01_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Sonar Sweep Part 1 with the real puzzle data
func TestSonarSweepPart1Puzzle(t *testing.T) {
	expected := 1553
	actual := SonarSweepPart1("../../../data/year2021/day_01_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Sonar Sweep Part 2 with the sample data
func TestSonarSweepPart2Sample(t *testing.T) {
	expected := 5
	actual := SonarSweepPart2("../../../data/year2021/day_01_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Sonar Sweep Part 2 with the real puzzle data
func TestSonarSweepPart2Puzzle(t *testing.T) {
	expected := 1597
	actual := SonarSweepPart2("../../../data/year2021/day_01_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
