package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Syntax Scoring Part 1 with the sample data
func TestSyntaxScoringPart1Sample(t *testing.T) {
	expected := 26397
	actual := SyntaxScoringPart1("../../../data/year2021/day_10_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Syntax Scoring Part 1 with the real puzzle data
func TestSyntaxScoringPart1Puzzle(t *testing.T) {
	expected := 367227
	actual := SyntaxScoringPart1("../../../data/year2021/day_10_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Syntax Scoring Part 2 with the sample data
func TestSyntaxScoringPart2Sample(t *testing.T) {
	expected := 288957
	actual := SyntaxScoringPart2("../../../data/year2021/day_10_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Syntax Scoring Part 2 with the real puzzle data
func TestSyntaxScoringPart2Puzzle(t *testing.T) {
	expected := 3583341858
	actual := SyntaxScoringPart2("../../../data/year2021/day_10_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
