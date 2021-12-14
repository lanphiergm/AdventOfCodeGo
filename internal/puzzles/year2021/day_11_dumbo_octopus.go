package year2021

import (
	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Dumbo Octopus Part 1 computes the number of flashes that have occurred after 100 steps
func DumboOctopusPart1(filename string) interface{} {
	grid := parseGrid(filename)
	flashCount := 0
	for i := 0; i < 100; i++ {
		flashCount += incrementGrid(&grid)
	}
	return flashCount
}

// Dumbo Octopus Part 2 computes the first iteration where all octopodes flash
// at the same time
func DumboOctopusPart2(filename string) interface{} {
	grid := parseGrid(filename)
	i := 0
	for {
		i++
		if incrementGrid(&grid) == 100 {
			return i
		}
	}
}

func incrementGrid(grid *[10][10]int) int {
	var hasFlashed [10][10]bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			(*grid)[i][j]++
			if (*grid)[i][j] > 9 {
				flashOctopus(grid, &hasFlashed, i, j)
			}
		}
	}

	hasChanged := true
	for hasChanged {
		hasChanged = false
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (*grid)[i][j] > 9 && !hasFlashed[i][j] {
					flashOctopus(grid, &hasFlashed, i, j)
					hasChanged = true
				}
			}
		}
	}

	flashCount := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (*grid)[i][j] > 9 {
				(*grid)[i][j] = 0
				flashCount++
			}
		}
	}
	return flashCount
}

func flashOctopus(grid *[10][10]int, hasFlashed *[10][10]bool, y int, x int) {
	(*hasFlashed)[y][x] = true
	minX := x - 1
	maxX := x + 1
	minY := y - 1
	maxY := y + 1
	if minX < 0 {
		minX = 0
	}
	if maxX > 9 {
		maxX = 9
	}
	if minY < 0 {
		minY = 0
	}
	if maxY > 9 {
		maxY = 9
	}
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if i != y || j != x {
				(*grid)[i][j]++
			}
		}
	}
}

func parseGrid(filename string) [10][10]int {
	data := utils.ReadStrings(filename)
	var grid [10][10]int
	for i, row := range data {
		for j, cell := range row {
			grid[i][j] = utils.Rtoi(cell)
		}
	}
	return grid
}
