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
	distances := make(map[utils.Coord]int)
	distances[utils.Coord{X: 0, Y: 0}] = 0
	hasVisited := make(map[utils.Coord]bool)
	queue := make([]node, 0)
	queue = enqueueCoord(queue, utils.Coord{X: 0, Y: 0}, 0)
	for len(queue) > 0 {
		src := queue[0].Coord
		queue = queue[1:]
		if hasVisited[src] {
			continue
		}
		hasVisited[src] = true
		adjacents := make([]utils.Coord, 0)
		if src.X > 0 { // left
			adjacents = append(adjacents, utils.Coord{X: src.X - 1, Y: src.Y})
		}
		if src.X < len((*grid)[0])-1 { // right
			adjacents = append(adjacents, utils.Coord{X: src.X + 1, Y: src.Y})
		}
		if src.Y > 0 { // up
			adjacents = append(adjacents, utils.Coord{X: src.X, Y: src.Y - 1})
		}
		if src.Y < len(*grid)-1 { // down
			adjacents = append(adjacents, utils.Coord{X: src.X, Y: src.Y + 1})
		}
		for _, dest := range adjacents {
			if hasVisited[dest] {
				continue
			}
			dist := distances[src] + (*grid)[dest.Y][dest.X]
			if existingDist, found := distances[dest]; !found || dist < existingDist {
				distances[dest] = dist
				queue = enqueueCoord(queue, dest, dist)
			}
		}
	}
	return distances[utils.Coord{X: len((*grid)[0]) - 1, Y: len(*grid) - 1}]
}

type node struct {
	Coord utils.Coord
	Dist  int
}

func enqueueCoord(queue []node, coord utils.Coord, dist int) []node {
	newNode := node{Coord: coord, Dist: dist}
	insertAt := len(queue)
	for index, item := range queue {
		if item.Dist > dist {
			insertAt = index
			break
		}
	}

	if len(queue) == insertAt {
		return append(queue, newNode)
	}
	queue = append(queue[:insertAt+1], queue[insertAt:]...)
	queue[insertAt] = newNode
	return queue
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
