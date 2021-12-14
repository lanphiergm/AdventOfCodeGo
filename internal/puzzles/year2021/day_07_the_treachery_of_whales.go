package year2021

import (
	"math"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// The Treachery of Whales Part 1 computes the minimum fuel consumption assuming a
// constant cost function
func TheTreacheryofWhalesPart1(filename string) interface{} {
	return findMinFuelConsumption(filename, sumLinearFuelConsumption)
}

// The Treachery of Whales Part 2 computes the minimum fuel consumption assuming an
// increasing cost function
func TheTreacheryofWhalesPart2(filename string) interface{} {
	return findMinFuelConsumption(filename, sumIncreasingFuelConsumption)
}

func findMinFuelConsumption(filename string, sumFuel sumFuelConsumptionFunc) int {
	positions := utils.ReadCsv(filename)
	min, max := utils.GetExtrema(positions)
	minFuel := utils.MaxInt
	for i := min; i <= max; i++ {
		fuel := sumFuel(positions, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

// Define the functions for computing fuel consumption
type sumFuelConsumptionFunc func(positions []int, rallyPoint int) int

func sumLinearFuelConsumption(positions []int, rallyPoint int) int {
	fuel := 0
	for i := 0; i < len(positions); i++ {
		fuel += int(math.Abs(float64(positions[i] - rallyPoint)))
	}
	return fuel
}

func sumIncreasingFuelConsumption(positions []int, rallyPoint int) int {
	fuel := 0
	for i := 0; i < len(positions); i++ {
		fuel += utils.Summorial(int(math.Abs(float64(positions[i] - rallyPoint))))
	}
	return fuel
}
