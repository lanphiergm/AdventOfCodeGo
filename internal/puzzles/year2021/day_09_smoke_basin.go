package year2021

import (
	"sort"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Smoke Basin Part 1 computes the sum of all risk levels for the heightmap
func SmokeBasinPart1(filename string) interface{} {
	data := utils.ReadStrings(filename)
	sum := 0
	for i, row := range data {
		for j, cell := range row {
			v := utils.Rtoi(cell)
			if i > 0 && v >= utils.Btoi(data[i-1][j]) { // up
				continue
			}
			if i < len(data)-1 && v >= utils.Btoi(data[i+1][j]) { // down
				continue
			}
			if j > 0 && v >= utils.Btoi(row[j-1]) { // left
				continue
			}
			if j < len(row)-1 && v >= utils.Btoi(row[j+1]) {
				continue
			}
			sum += v + 1
		}
	}
	return sum
}

// Smoke Basin Part 2 computes the product of the size of the three largest basins
func SmokeBasinPart2(filename string) interface{} {
	data := utils.ReadStrings(filename)
	basinSizes := []int{}
	for i, row := range data {
		for j, cell := range row {
			v := utils.Rtoi(cell)
			if i > 0 && v >= utils.Btoi(data[i-1][j]) { // up
				continue
			}
			if i < len(data)-1 && v >= utils.Btoi(data[i+1][j]) { // down
				continue
			}
			if j > 0 && v >= utils.Btoi(row[j-1]) { // left
				continue
			}
			if j < len(row)-1 && v >= utils.Btoi(row[j+1]) {
				continue
			}
			coords := map[coord]struct{}{}
			findBasinCoords(data, coords, i, j)
			basinSizes = append(basinSizes, len(coords))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

type coord struct {
	x, y int
}

func findBasinCoords(data []string, coords map[coord]struct{}, initialI int, initialJ int) {
	if _, ok := coords[coord{initialI, initialJ}]; ok {
		return
	}
	coords[coord{initialI, initialJ}] = struct{}{}
	for i := initialI - 1; i >= 0; i-- { // up
		v := utils.Btoi(data[i][initialJ])
		if v == 9 {
			break
		}
		findBasinCoords(data, coords, i, initialJ)
	}
	for i := initialI + 1; i < len(data); i++ { // down
		v := utils.Btoi(data[i][initialJ])
		if v == 9 {
			break
		}
		findBasinCoords(data, coords, i, initialJ)
	}
	for j := initialJ - 1; j >= 0; j-- { // left
		v := utils.Btoi(data[initialI][j])
		if v == 9 {
			break
		}
		findBasinCoords(data, coords, initialI, j)
	}
	for j := initialJ + 1; j < len(data[initialI]); j++ { // right
		v := utils.Btoi(data[initialI][j])
		if v == 9 {
			break
		}
		findBasinCoords(data, coords, initialI, j)
	}
}
