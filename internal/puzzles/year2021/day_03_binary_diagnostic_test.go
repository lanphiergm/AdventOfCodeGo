package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/data/year2021data"
	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Binary Diagnostic Part 1 with the sample data
func TestBinaryDiagnosticPart1Sample(t *testing.T) {
	expected := 198
	actual := BinaryDiagnosticPart1(year2021data.BinaryDiagnosticSampleData)
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Binary Diagnostic Part 1 with the real puzzle data
func TestBinaryDiagnosticPart1Puzzle(t *testing.T) {
	expected := 3901196
	actual := BinaryDiagnosticPart1(year2021data.BinaryDiagnosticPuzzleData)
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Binary Diagnostic Part 2 with the sample data
func TestBinaryDiagnosticPart2Sample(t *testing.T) {
	expected := 230
	actual := BinaryDiagnosticPart2(year2021data.BinaryDiagnosticSampleData)
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Binary Diagnostic Part 2 with the real puzzle data
func TestBinaryDiagnosticPart2Puzzle(t *testing.T) {
	expected := 4412188
	actual := BinaryDiagnosticPart2(year2021data.BinaryDiagnosticPuzzleData)
	utils.AssertAreEqual(t, expected, actual)
}
