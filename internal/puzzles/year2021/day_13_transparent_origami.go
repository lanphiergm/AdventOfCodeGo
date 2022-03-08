package year2021

import (
	"fmt"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Transparent Origami Part 1 computes the number of visible dots after one fold
func TransparentOrigamiPart1(filename string) interface{} {
	coords, folds := parseOrigamiData(filename)
	coords = performFolds(coords, folds[:1])
	return len(coords)
}

// Transparent Origami Part 2 computes the activation code
func TransparentOrigamiPart2(filename string) interface{} {
	coords, folds := parseOrigamiData(filename)
	coords = performFolds(coords, folds)
	//printCoords(coords)
	return translateCoords(coords)
}

func performFolds(coords map[utils.Coord]struct{}, folds []string) map[utils.Coord]struct{} {
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

func performFoldX(coords map[utils.Coord]struct{}, foldAt int) map[utils.Coord]struct{} {
	for c := range coords {
		if c.X == foldAt {
			delete(coords, c)
		} else if c.X > foldAt {
			coords[utils.Coord{X: 2*foldAt - c.X, Y: c.Y}] = struct{}{}
			delete(coords, c)
		}
	}
	return coords
}

func performFoldY(coords map[utils.Coord]struct{}, foldAt int) map[utils.Coord]struct{} {
	for c := range coords {
		if c.Y == foldAt {
			delete(coords, c)
		} else if c.Y > foldAt {
			coords[utils.Coord{X: c.X, Y: 2*foldAt - c.Y}] = struct{}{}
			delete(coords, c)
		}
	}
	return coords
}

func printCoords(coords map[utils.Coord]struct{}) {
	maxX := 0
	maxY := 0
	for c := range coords {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	toPrint := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		toPrint[i] = make([]bool, maxX+1)
	}
	for c := range coords {
		toPrint[c.Y][c.X] = true
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

func translateCoords(coords map[utils.Coord]struct{}) string {
	var grid [6][39]bool
	for c := range coords {
		grid[c.Y][c.X] = true
	}

	code := ""
	for i := 0; i < 39; i += 5 {
		for key, val := range letters {
			isMatch := true
			for j := 0; j < 6; j++ {
				for k := 0; k < 4; k++ {
					if key[j][k] != grid[j][i+k] {
						isMatch = false
						break
					}
				}
				if !isMatch {
					break
				}
			}
			if isMatch {
				code += val
				break
			}
		}
	}

	return code
}

var letterE = [6][4]bool{{true, true, true, true}, {true, false, false, false}, {true, true, true, false},
	{true, false, false, false}, {true, false, false, false}, {true, true, true, true}}
var letterG = [6][4]bool{{false, true, true, false}, {true, false, false, true}, {true, false, false, false},
	{true, false, true, true}, {true, false, false, true}, {false, true, true, true}}
var letterH = [6][4]bool{{true, false, false, true}, {true, false, false, true}, {true, true, true, true},
	{true, false, false, true}, {true, false, false, true}, {true, false, false, true}}
var letterJ = [6][4]bool{{false, false, true, true}, {false, false, false, true}, {false, false, false, true},
	{false, false, false, true}, {true, false, false, true}, {false, true, true, false}}
var letterL = [6][4]bool{{true, false, false, false}, {true, false, false, false}, {true, false, false, false},
	{true, false, false, false}, {true, false, false, false}, {true, true, true, true}}
var letterU = [6][4]bool{{true, false, false, true}, {true, false, false, true}, {true, false, false, true},
	{true, false, false, true}, {true, false, false, true}, {false, true, true, false}}
var letters = map[[6][4]bool]string{letterE: "E", letterG: "G", letterH: "H", letterJ: "J", letterL: "L", letterU: "U"}

func parseOrigamiData(filename string) (map[utils.Coord]struct{}, []string) {
	data := utils.ReadStrings(filename)
	coords := make(map[utils.Coord]struct{})
	foldStart := 0
	for i, c := range data {
		if c == "" {
			foldStart = i + 1
			break
		}
		tmp := strings.Split(c, ",")
		coords[utils.Coord{X: utils.Atoi(tmp[0]), Y: utils.Atoi(tmp[1])}] = struct{}{}
	}
	return coords, data[foldStart:]
}
