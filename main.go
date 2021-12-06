// Main application package
package main

import (
	"fmt"

	"github.com/lanphiergm/adventofcodego/internal/data/year2021data"
	"github.com/lanphiergm/adventofcodego/internal/puzzles/year2021"
)

// Runs all puzzles
func main() {
	fmt.Printf("2021 Day 01 Part 1: %v\n",
		year2021.SonarSweepPart1(year2021data.SonarSweepPuzzleData))
	fmt.Printf("2021 Day 01 Part 2: %v\n",
		year2021.SonarSweepPart2(year2021data.SonarSweepPuzzleData))
	fmt.Printf("2021 Day 02 Part 1: %v\n",
		year2021.DivePart1(year2021data.DivePuzzleData))
	fmt.Printf("2021 Day 02 Part 2: %v\n",
		year2021.DivePart2(year2021data.DivePuzzleData))
	fmt.Printf("2021 Day 03 Part 1: %v\n",
		year2021.BinaryDiagnosticPart1(year2021data.BinaryDiagnosticPuzzleData))
	fmt.Printf("2021 Day 03 Part 2: %v\n",
		year2021.BinaryDiagnosticPart2(year2021data.BinaryDiagnosticPuzzleData))
}
