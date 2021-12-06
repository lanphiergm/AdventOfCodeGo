package year2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Dive! Part 1 computes the horizontal position and depth after a series of
// commands are executed
func DivePart1(filename string) int {
	hPos := 0
	depth := 0

	commands := utils.ReadStrings(filename)
	for _, command := range commands {
		s := strings.Split(command, " ")
		v, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		if s[0] == "forward" {
			hPos += v
		} else if s[0] == "down" {
			depth += v
		} else if s[0] == "up" {
			depth -= v
		} else {
			fmt.Printf("Unknown command: %v", s[0])
			return 0
		}
	}

	return hPos * depth
}

// Dive! Part 2 computes the horizontal position and depth after a series of
// commands are executed with a different understanding of the command syntax
func DivePart2(filename string) int {
	hPos := 0
	depth := 0
	aim := 0

	commands := utils.ReadStrings(filename)
	for _, command := range commands {
		s := strings.Split(command, " ")
		v, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		if s[0] == "forward" {
			hPos += v
			depth += v * aim
		} else if s[0] == "down" {
			aim += v
		} else if s[0] == "up" {
			aim -= v
		} else {
			fmt.Printf("Unknown command: %v", s[0])
			return 0
		}
	}
	return hPos * depth
}
