package year2021

import (
	"log"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Snailfish Part 1 computes the magnitude of the sum of the numbers
func SnailfishPart1(filename string) interface{} {
	data := utils.ReadStrings(filename)
	var pair *snailfishPair
	for _, s := range data {
		if pair == nil {
			pair = parseSnailfishPair(s)
		} else {
			pair = &snailfishPair{Left: pair, Right: parseSnailfishPair(s)}
			pair.Left.Parent = pair
			pair.Right.Parent = pair
			pair.reduce()
		}
	}
	return pair.magnitude()
}

// Snailfish Part 2 computes the maximum sum of any two numbers
func SnailfishPart2(filename string) interface{} {
	data := utils.ReadStrings(filename)
	maxMagnitude := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if i == j {
				continue
			}
			// We need to parse them fresh each time because subsequent
			// operations will change cached structs
			pair := &snailfishPair{}
			pair.Left = parseSnailfishPair(data[i])
			pair.Left.Parent = pair
			pair.Right = parseSnailfishPair(data[j])
			pair.Right.Parent = pair
			pair.reduce()
			magnitude := pair.magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}
	return maxMagnitude
}

type snailfishPair struct {
	Number int
	Parent *snailfishPair
	Left   *snailfishPair
	Right  *snailfishPair
}

func (p snailfishPair) depth() int {
	if p.Parent == nil {
		return 1
	} else {
		return p.Parent.depth() + 1
	}
}

func (p snailfishPair) isNumber() bool {
	return p.Left == nil && p.Right == nil
}

func (p snailfishPair) magnitude() int {
	if p.isNumber() {
		return p.Number
	} else {
		return 3*p.Left.magnitude() + 2*p.Right.magnitude()
	}
}

func (p *snailfishPair) reduce() {
	isReduced := false
	for !isReduced {
		if p.explode() {
			continue
		} else if p.split() {
			continue
		}
		isReduced = true
	}
}

func (p *snailfishPair) explode() bool {
	if p.isNumber() {
		return false
	} else if p.depth() > 4 && p.Left.isNumber() && p.Right.isNumber() {
		prev := p.prev()
		if prev != nil {
			*prev = *prev + p.Left.Number
		}
		next := p.next()
		if next != nil {
			*next = *next + p.Right.Number
		}
		p.Number = 0
		p.Left = nil
		p.Right = nil
		return true
	} else {
		hasExploded := p.Left.explode()
		if !hasExploded {
			hasExploded = p.Right.explode()
		}
		return hasExploded
	}
}

func (p *snailfishPair) prev() *int {
	var val *int
	for p.Parent != nil && p.Parent.Left == p {
		p = p.Parent
	}
	if p.Parent != nil && p.Parent.Left != p {
		p = p.Parent.Left
		for !p.isNumber() {
			p = p.Right
		}
		val = &p.Number
	}
	return val
}

func (p *snailfishPair) next() *int {
	var val *int
	for p.Parent != nil && p.Parent.Right == p {
		p = p.Parent
	}
	if p.Parent != nil && p.Parent.Right != p {
		p = p.Parent.Right
		for !p.isNumber() {
			p = p.Left
		}
		val = &p.Number
	}
	return val
}

func (p *snailfishPair) split() bool {
	if p.isNumber() && p.Number > 9 {
		left := new(snailfishPair)
		left.Number = p.Number / 2
		left.Parent = p
		right := new(snailfishPair)
		if p.Number%2 == 0 {
			right.Number = p.Number / 2
		} else {
			right.Number = p.Number/2 + 1
		}
		right.Parent = p
		p.Number = 0
		p.Left = left
		p.Right = right
		return true
	} else if p.isNumber() {
		return false
	} else {
		hasSplit := p.Left.split()
		if !hasSplit {
			hasSplit = p.Right.split()
		}
		return hasSplit
	}
}

func parseSnailfishPair(s string) *snailfishPair {
	if s[0] == '[' {
		bracketCount := 0
		for i := 1; i < len(s); i++ {
			if s[i] == '[' {
				bracketCount++
			} else if s[i] == ']' {
				bracketCount--
			}
			if bracketCount == 0 {
				left := parseSnailfishPair(s[1 : i+1])
				right := parseSnailfishPair(s[i+2 : len(s)-1])
				p := new(snailfishPair)
				p.Left = left
				p.Right = right
				left.Parent = p
				right.Parent = p
				return p
			}
		}
	} else {
		p := new(snailfishPair)
		p.Number = utils.Atoi(s)
		return p
	}

	log.Fatal("Parse Error")
	return &snailfishPair{}
}
