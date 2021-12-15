package year2021

import (
	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Chiton Part 1 computes the lowest total risk level for a path through the cave
func ChitonPart1(filename string) interface{} {
	grid := parseChitonGrid(filename)
	return relaxGrid(&grid)
}

// Chiton Part 2 computes the lowest total risk level for a path through the expanded cave
func ChitonPart2(filename string) interface{} {
	grid := parseChitonGrid(filename)
	grid = replicateGrid(&grid)
	return relaxGrid(&grid)
}

func relaxGrid(grid *[][]int) interface{} {
	lenX := len((*grid)[0])
	lenY := len(*grid)
	distances := make(map[coord]int)
	distances[coord{0, 0}] = 0
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for i := 0; i < lenY; i++ {
			for j := 0; j < lenX; j++ {
				hasChanged = relaxAdjacent(grid, &distances, coord{j, i}) || hasChanged
			}
		}
	}
	return distances[coord{lenX - 1, lenY - 1}]
}

func relaxAdjacent(grid *[][]int, distances *map[coord]int, src coord) bool {
	hasChanged := false
	coords := make([]coord, 0)
	if src.x > 0 { // left
		coords = append(coords, coord{src.x - 1, src.y})
	}
	if src.x < len((*grid)[0])-1 { // right
		coords = append(coords, coord{src.x + 1, src.y})
	}
	if src.y > 0 { // up
		coords = append(coords, coord{src.x, src.y - 1})
	}
	if src.y < len(*grid)-1 { // down
		coords = append(coords, coord{src.x, src.y + 1})
	}
	for _, dest := range coords {
		dist := (*distances)[src] + (*grid)[dest.y][dest.x]
		if existingDist, found := (*distances)[dest]; !found || dist < existingDist {
			(*distances)[dest] = dist
			hasChanged = true
		}
	}
	return hasChanged
}

func parseChitonGrid(filename string) [][]int {
	data := utils.ReadStrings(filename)
	grid := make([][]int, len(data))
	for i, row := range data {
		grid[i] = make([]int, len(row))
		for j, r := range row {
			grid[i][j] = utils.Rtoi(r)
		}
	}
	return grid
}

func replicateGrid(grid *[][]int) [][]int {
	lenX := len((*grid)[0])
	lenY := len(*grid)
	newGrid := make([][]int, lenY*5)
	for i := 0; i < lenY*5; i++ {
		newGrid[i] = make([]int, lenX*5)
	}
	for i := 0; i < 5; i++ {
		for y := 0; y < lenY; y++ {
			for x := 0; x < lenX; x++ {
				if i == 0 { // copy from original grid
					newGrid[y][x] = (*grid)[y][x]
				} else { // copy down from above
					newV := newGrid[(i-1)*lenY+y][x] + 1
					if newV == 10 {
						newV = 1
					}
					newGrid[i*lenY+y][x] = newV

				}
			}
		}

		// copy right
		for j := 1; j < 5; j++ {
			for y := 0; y < lenY; y++ {
				for x := 0; x < lenX; x++ {
					newV := newGrid[i*lenY+y][(j-1)*lenX+x] + 1
					if newV == 10 {
						newV = 1
					}
					newGrid[i*lenY+y][j*lenX+x] = newV
				}
			}
		}
	}
	return newGrid
}
