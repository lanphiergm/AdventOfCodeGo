// Main application package
package main

import (
	"fmt"
	"time"

	"github.com/lanphiergm/adventofcodego/internal/puzzles/year2021"
	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Runs all puzzles
func main() {
	runPuzzle(2021, 1, 1, year2021.SonarSweepPart1)
	runPuzzle(2021, 1, 2, year2021.SonarSweepPart2)
	runPuzzle(2021, 2, 1, year2021.DivePart1)
	runPuzzle(2021, 2, 2, year2021.DivePart2)
	runPuzzle(2021, 3, 1, year2021.BinaryDiagnosticPart1)
	runPuzzle(2021, 3, 2, year2021.BinaryDiagnosticPart2)
	runPuzzle(2021, 4, 1, year2021.GiantSquidPart1)
	runPuzzle(2021, 4, 2, year2021.GiantSquidPart2)
	runPuzzle(2021, 5, 1, year2021.HydrothermalVenturePart1)
	runPuzzle(2021, 5, 2, year2021.HydrothermalVenturePart2)
	runPuzzle(2021, 6, 1, year2021.LanternfishPart1)
	runPuzzle(2021, 6, 2, year2021.LanternfishPart2)
	runPuzzle(2021, 7, 1, year2021.TheTreacheryofWhalesPart1)
	runPuzzle(2021, 7, 2, year2021.TheTreacheryofWhalesPart2)

	min, mean, max := computeStats()
	fmt.Printf("===============================================================\n")
	fmt.Printf(" Puzzle count: %v\n", len(durations))
	fmt.Printf(" Min duration: %.4fms\n", float32(min)/1000.0)
	fmt.Printf("Mean duration: %.4fms\n", float32(mean)/1000.0)
	fmt.Printf(" Max duration: %.4fms\n", float32(max)/1000.0)
}

type puzzleFunc func(string) int

func runPuzzle(year int, day int, part int, puzzle puzzleFunc) {
	filename := fmt.Sprintf("data/year%v/day_%02v_puzzle_data.txt", year, day)
	start := time.Now()
	result := puzzle(filename)
	elapsed := time.Since(start)
	fmt.Printf("%v Day %02v Part %v: %-20v (%v)\n", year, day, part, result, elapsed)
	durations = append(durations, elapsed)
}

var durations []time.Duration

func computeStats() (int, int, int) {
	min := utils.MaxInt
	sum := 0
	max := utils.MinInt
	for _, e := range durations {
		v := int(e.Microseconds())
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, sum / len(durations), max
}
