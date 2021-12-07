package year2021

import (
	"math"
	"regexp"
	"strconv"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// HydrothermalVenture Part 1 computes {part 1 description}
func HydrothermalVenturePart1(filename string) int {
	return findDangerousAreas(filename, false)
}

// HydrothermalVenture Part 2 computes {part 2 description}
func HydrothermalVenturePart2(filename string) int {
	return findDangerousAreas(filename, true)
}

func keepGoingAscending(counter int, end int) bool {
	return counter <= end
}
func keepGoingDescending(counter int, end int) bool {
	return counter >= end
}

func findDangerousAreas(filename string, includeDiagonals bool) int {
	lines := utils.ReadStrings(filename)
	var grid [1000][1000]int
	rs := `(?P<x1>\d{1,3}),(?P<y1>\d{1,3}) -> (?P<x2>\d{1,3}),(?P<y2>\d{1,3})`
	r := regexp.MustCompile(rs)
	for _, line := range lines {
		x1, y1, x2, y2 := parseLine(r, line)
		if x1 == x2 || y1 == y2 || (includeDiagonals && isDiagonal(x1, y1, x2, y2)) {

			// Set up how the x variable will proceed through the loop
			xInc := 1
			xKeepGoing := keepGoingAscending
			if x2 < x1 {
				xInc = -1
				xKeepGoing = keepGoingDescending
			} else if x1 == x2 {
				xInc = 0
			}

			// Set up how the y variable will proceed through the loop
			yInc := 1
			yKeepGoing := keepGoingAscending
			if y2 < y1 {
				yInc = -1
				yKeepGoing = keepGoingDescending
			} else if y1 == y2 {
				yInc = 0
			}

			// Perform the loop
			for i, j := y1, x1; yKeepGoing(i, y2) && xKeepGoing(j, x2); i, j = i+yInc, j+xInc {
				grid[i][j]++
			}
		}
	}

	dangerousAreas := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				dangerousAreas++
			}
		}
	}

	return dangerousAreas
}

func isDiagonal(x1 int, y1 int, x2 int, y2 int) bool {
	slope := float64(y2-y1) / float64(x2-x1)
	return math.Abs(slope) == 1.0
}

func parseLine(r *regexp.Regexp, line string) (int, int, int, int) {
	result := r.FindAllStringSubmatch(line, -1)
	x1, _ := strconv.ParseInt(result[0][1], 10, 32)
	y1, _ := strconv.ParseInt(result[0][2], 10, 32)
	x2, _ := strconv.ParseInt(result[0][3], 10, 32)
	y2, _ := strconv.ParseInt(result[0][4], 10, 32)
	return int(x1), int(y1), int(x2), int(y2)
}
