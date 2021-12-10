package year2021

import (
	"sort"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Syntax Scoring Part 1 computes the syntax error score
func SyntaxScoringPart1(filename string) int {
	lines := utils.ReadStrings(filename)
	score := 0
	points := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, line := range lines {
		if doScore, b := isCorrupted(line); doScore {
			score += points[b]
		}
	}
	return score
}

// Syntax Scoring Part 2 computes the middle autocomplete score
func SyntaxScoringPart2(filename string) int {
	lines := utils.ReadStrings(filename)
	scores := []int{}
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	for _, line := range lines {
		if doScore, completion := isIncomplete(line); doScore {
			score := 0
			for _, r := range completion {
				score *= 5
				score += points[r]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func isCorrupted(line string) (bool, byte) {
	_, _, isErr, errByte := matchChunk(line, 0)
	return isErr, errByte
}

func isIncomplete(line string) (bool, string) {
	_, isEof, _, completion := completeChunk(line, 0)
	return isEof, completion
}

var openers = []byte{'(', '[', '{', '<'}
var closures = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func isOpener(test byte) bool {
	for _, b := range openers {
		if test == b {
			return true
		}
	}
	return false
}

func matchChunk(line string, pos int) (int, bool, bool, byte) {
	if pos == len(line)-1 {
		return 0, true, false, 0
	}

	i := pos
	for isOpener(line[i+1]) {
		count, isEof, isErr, errByte := matchChunk(line, i+1)
		if isEof {
			return 0, true, false, 0
		} else if isErr {
			return 0, false, true, errByte
		} else {
			i += count
		}

		if i == len(line)-1 {
			return 0, true, false, 0
		}
	}

	if line[i+1] == closures[line[pos]] {
		return 2 + i - pos, false, false, 0
	} else {
		return 0, false, true, line[i+1]
	}
}

func completeChunk(line string, pos int) (int, bool, bool, string) {
	if pos == len(line)-1 {
		return 1, true, false, string(closures[line[pos]])
	}

	completion := ""
	i := pos
	for isOpener(line[i+1]) {
		count, isEof, isErr, c := completeChunk(line, i+1)
		i += count
		if isEof {
			completion += c
		} else if isErr {
			return 0, false, true, ""
		}

		if i == len(line)-1 {
			return 1 + i - pos, true, false, completion + string(closures[line[pos]])
		}
	}

	if line[i+1] == closures[line[pos]] {
		return 2 + i - pos, false, false, ""
	} else {
		return 0, false, true, ""
	}
}
