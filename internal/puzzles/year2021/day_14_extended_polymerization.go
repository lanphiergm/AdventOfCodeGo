package year2021

import (
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Extended Polymerization Part 1 computes the difference between the most and least common
// elements after 10 steps
func ExtendedPolymerizationPart1(filename string) interface{} {
	return insertPolymers(filename, 10)
}

// Extended Polymerization Part 2 computes the difference between the most and least common
// elements after 40 steps
func ExtendedPolymerizationPart2(filename string) interface{} {
	return insertPolymers(filename, 40)
}

func insertPolymers(filename string, iterations int) interface{} {
	p, insertions := parsePolymerData(filename)
	counts := make(map[string]int)
	for i := 0; i < len(p); i++ {
		counts[p[i:i+1]]++
	}
	pairCounts := make(map[string]int)
	for i := 0; i < len(p)-1; i++ {
		pairCounts[p[i:i+2]]++
	}

	for n := 0; n < iterations; n++ {
		newPairCounts := deepCopyMap(pairCounts)
		for k, v := range insertions {
			if c, found := pairCounts[k]; found {
				newPairCounts[k] -= c
				counts[v] += c
				left := k[0:1] + v
				right := v + k[1:2]
				newPairCounts[left] += c
				newPairCounts[right] += c
			}
		}
		pairCounts = newPairCounts
	}

	max := utils.MinInt
	min := utils.MaxInt
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func deepCopyMap(originalMap map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}

func parsePolymerData(filename string) (string, map[string]string) {
	data := utils.ReadStrings(filename)
	insertions := make(map[string]string)

	for i := 2; i < len(data); i++ {
		tmp := strings.Split(data[i], " -> ")
		insertions[tmp[0]] = tmp[1]
	}

	return data[0], insertions
}
