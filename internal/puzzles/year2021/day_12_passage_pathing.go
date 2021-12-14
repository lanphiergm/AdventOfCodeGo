package year2021

import (
	"fmt"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Passage Pathing Part 1 computes the number of paths through the cave system if
// small caves can be visited only once
func PassagePathingPart1(filename string) interface{} {
	dsts := parsePaths(filename)
	paths := findPathsTo(dsts, []string{",end"}, "end")
	return len(paths)
}

// Passage Pathing Part 2 computes the number of paths through the cave system if
// a single small cave can be visited twice
func PassagePathingPart2(filename string) interface{} {
	dsts := parsePaths(filename)
	paths := findPathsToWithDoubleSmall(dsts, []string{",end"}, "end", false)
	return len(paths)
}

// Depth-first search starting at the end and working backwards to the start
func findPathsTo(dsts map[string][]string, paths []string, dst string) []string {
	srcs := dsts[dst]
	newPaths := make([]string, 0)
	for _, src := range srcs {
		for _, path := range paths {
			if src == "start" {
				newPaths = append(newPaths, fmt.Sprintf("start%v", path))
			} else if strings.ToUpper(src) == src || !strings.Contains(path, fmt.Sprintf(",%v,", src)) {
				newPaths = append(newPaths, findPathsTo(dsts, []string{fmt.Sprintf(",%v%v", src, path)}, src)...)
			}
		}
	}
	return newPaths
}

// Depth-first search starting at the end and working backwards to the start
func findPathsToWithDoubleSmall(dsts map[string][]string, paths []string, dst string, hasDoubleSmall bool) []string {
	srcs := dsts[dst]
	newPaths := make([]string, 0)
	for _, src := range srcs {
		for _, path := range paths {
			if src == "start" {
				newPaths = append(newPaths, fmt.Sprintf("start%v", path))
			} else if strings.ToUpper(src) == src {
				newPaths = append(newPaths, findPathsToWithDoubleSmall(dsts, []string{fmt.Sprintf(",%v%v", src, path)}, src, hasDoubleSmall)...)
			} else if !strings.Contains(path, fmt.Sprintf(",%v,", src)) {
				newPaths = append(newPaths, findPathsToWithDoubleSmall(dsts, []string{fmt.Sprintf(",%v%v", src, path)}, src, hasDoubleSmall)...)
			} else if !hasDoubleSmall {
				newPaths = append(newPaths, findPathsToWithDoubleSmall(dsts, []string{fmt.Sprintf(",%v%v", src, path)}, src, true)...)
			}
		}
	}
	return newPaths
}

func parsePaths(filename string) map[string][]string {
	data := utils.ReadStrings(filename)
	dsts := make(map[string][]string)
	for _, line := range data {
		pair := strings.Split(line, "-")

		// first direction
		srcs := dsts[pair[1]]
		if srcs == nil {
			srcs = make([]string, 0)
		}
		if pair[0] != "end" && pair[1] != "start" {
			srcs = append(srcs, pair[0])
			dsts[pair[1]] = srcs
		}

		// second direction
		srcs = dsts[pair[0]]
		if srcs == nil {
			srcs = make([]string, 0)
		}
		if pair[1] != "end" && pair[0] != "start" {
			srcs = append(srcs, pair[1])
			dsts[pair[0]] = srcs
		}
	}
	return dsts
}
