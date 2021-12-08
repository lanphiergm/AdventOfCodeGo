package year2021

import (
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// GiantSquid Part 1 computes the score of the first winning board
func GiantSquidPart1(filename string) int {
	draws, boards := parseData(filename)
	drawMap := make(map[int]struct{})
	for i := range draws {
		drawMap[draws[i]] = struct{}{}
		if i <= 4 {
			continue
		} // we can't have a winner with fewer than 5 draws
		for j := range boards {
			won, score := isWinner(boards[j], drawMap, draws[i])
			if won {
				return score
			}
		}
	}

	return 0
}

// GiantSquid Part 2 computes the score of the last winning board
func GiantSquidPart2(filename string) int {
	draws, boards := parseData(filename)
	drawMap := make(map[int]struct{})
	winners := make(map[int]struct{})
	for i := range draws {
		drawMap[draws[i]] = struct{}{}
		if i <= 4 {
			// we can't have a winner with fewer than 5 draws
			continue
		}
		for j := range boards {
			_, alreadyWon := winners[j]
			if !alreadyWon {
				won, score := isWinner(boards[j], drawMap, draws[i])
				if won {
					winners[j] = struct{}{}
					if len(winners) == len(boards) {
						return score
					}
				}
			}
		}
	}

	return 0
}

func isWinner(board [][]int, draws map[int]struct{}, lastDraw int) (bool, int) {
	// By row first
	for i := 0; i < 5; i++ {
		isComplete := true
		for j := 0; j < 5; j++ {
			_, isFound := draws[board[i][j]]
			if !isFound {
				isComplete = false
				break
			}
		}
		if isComplete {
			return true, scoreBoard(board, draws, lastDraw)
		}
	}

	// By column next
	for i := 0; i < 5; i++ {
		isComplete := true
		for j := 0; j < 5; j++ {
			_, isFound := draws[board[j][i]]
			if !isFound {
				isComplete = false
				break
			}
		}
		if isComplete {
			return true, scoreBoard(board, draws, lastDraw)
		}
	}

	return false, 0
}

func scoreBoard(board [][]int, draws map[int]struct{}, lastDraw int) int {
	unmarkedSum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_, isFound := draws[board[i][j]]
			if !isFound {
				unmarkedSum += board[i][j]
			}
		}
	}

	return unmarkedSum * lastDraw
}

func parseData(filename string) ([]int, [][][]int) {
	lines := utils.ReadStrings(filename)
	draws := utils.ParseInts(lines[0])

	numBoards := (len(lines) - 1) / 6
	boards := make([][][]int, numBoards)
	for i := range boards {
		boards[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			boards[i][j] = make([]int, 5)
			cols := strings.Split(lines[2+6*i+j], " ")
			k := 0
			for _, val := range cols {
				if val != "" {
					boards[i][j][k] = utils.Atoi(val)
					k++
				}
			}
		}
	}

	return draws, boards
}
