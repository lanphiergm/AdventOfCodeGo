package year2021

import (
	"fmt"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Transparent Origami Part 1 computes the number of visible dots after one fold
func TransparentOrigamiPart1(filename string) int {
	coords, folds := parseOrigamiData(filename)
	coords = performFolds(coords, folds[:1])
	return len(coords)
}

// Transparent Origami Part 2 computes the activation code
func TransparentOrigamiPart2(filename string) int {
	coords, folds := parseOrigamiData(filename)
	coords = performFolds(coords, folds)
	printCoords(coords)
	return len(coords)
}

func performFolds(coords map[coord]struct{}, folds []string) map[coord]struct{} {
	foldAt := utils.Atoi(strings.Split(folds[0], "=")[1])
	if strings.HasPrefix(folds[0], "fold along x=") {
		coords = performFoldX(coords, foldAt)
	} else {
		coords = performFoldY(coords, foldAt)
	}
	if len(folds) > 1 {
		coords = performFolds(coords, folds[1:])
	}
	return coords
}

func performFoldX(coords map[coord]struct{}, foldAt int) map[coord]struct{} {
	for c := range coords {
		if c.x == foldAt {
			delete(coords, c)
		} else if c.x > foldAt {
			coords[coord{2*foldAt - c.x, c.y}] = struct{}{}
			delete(coords, c)
		}
	}
	return coords
}

func performFoldY(coords map[coord]struct{}, foldAt int) map[coord]struct{} {
	for c := range coords {
		if c.y == foldAt {
			delete(coords, c)
		} else if c.y > foldAt {
			coords[coord{c.x, 2*foldAt - c.y}] = struct{}{}
			delete(coords, c)
		}
	}
	return coords
}

func printCoords(coords map[coord]struct{}) {
	maxX := 0
	maxY := 0
	for c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	toPrint := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		toPrint[i] = make([]bool, maxX+1)
	}
	for c := range coords {
		toPrint[c.y][c.x] = true
	}
	for _, row := range toPrint {
		for _, v := range row {
			if v {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func parseOrigamiData(filename string) (map[coord]struct{}, []string) {
	data := utils.ReadStrings(filename)
	coords := make(map[coord]struct{})
	foldStart := 0
	for i, c := range data {
		if c == "" {
			foldStart = i + 1
			break
		}
		tmp := strings.Split(c, ",")
		coords[coord{utils.Atoi(tmp[0]), utils.Atoi(tmp[1])}] = struct{}{}
	}
	return coords, data[foldStart:]
}
