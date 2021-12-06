// Main application package
package main

import (
	"fmt"

	"github.com/lanphiergm/adventofcodego/internal/puzzles/year2021"
)

// Runs all puzzles
func main() {
	fmt.Printf("2021 Day 01 Part 1: %v\n",
		year2021.SonarSweepPart1("data/year2021/day_01_puzzle_data.txt"))
	fmt.Printf("2021 Day 01 Part 2: %v\n",
		year2021.SonarSweepPart2("data/year2021/day_01_puzzle_data.txt"))
	fmt.Printf("2021 Day 02 Part 1: %v\n",
		year2021.DivePart1("data/year2021/day_02_puzzle_data.txt"))
	fmt.Printf("2021 Day 02 Part 2: %v\n",
		year2021.DivePart2("data/year2021/day_02_puzzle_data.txt"))
	fmt.Printf("2021 Day 03 Part 1: %v\n",
		year2021.BinaryDiagnosticPart1("data/year2021/day_03_puzzle_data.txt"))
	fmt.Printf("2021 Day 03 Part 2: %v\n",
		year2021.BinaryDiagnosticPart2("data/year2021/day_03_puzzle_data.txt"))
	fmt.Printf("2021 Day 04 Part 1: %v\n",
		year2021.GiantSquidPart1("data/year2021/day_04_puzzle_data.txt"))
	fmt.Printf("2021 Day 04 Part 2: %v\n",
		year2021.GiantSquidPart2("data/year2021/day_04_puzzle_data.txt"))
	fmt.Printf("2021 Day 05 Part 1: %v\n",
		year2021.HydrothermalVenturePart1("data/year2021/day_05_puzzle_data.txt"))
	fmt.Printf("2021 Day 05 Part 2: %v\n",
		year2021.HydrothermalVenturePart2("data/year2021/day_05_puzzle_data.txt"))
}
