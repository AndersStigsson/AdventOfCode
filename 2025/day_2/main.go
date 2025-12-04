package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Solution struct {
	Ranges     []Range
	InvalidIds []int
}

type Range struct {
	Min    int
	Max    int
	MinStr string
	MaxStr string
}

func main() {
	solvePartOne()
	solvePartTwo()
}

func (s *Solution) parseInput(input string) {
	input = strings.TrimRight(input, "\n")
	vals := strings.SplitSeq(input, ",")
	for val := range vals {
		p := Range{}
		splitVal := strings.Split(val, "-")
		p.MinStr = splitVal[0]
		p.MaxStr = splitVal[1]
		v1, _ := strconv.Atoi(splitVal[0])
		v2, _ := strconv.Atoi(splitVal[1])
		p.Min = v1
		p.Max = v2
		s.Ranges = append(s.Ranges, p)

	}

}

func (s *Solution) solvePartOne() {
	for _, r := range s.Ranges {
		if len(r.MinStr)%2 != 0 && len(r.MinStr) == len(r.MaxStr) {
			continue
		}
		x := r.Min
		for x <= r.Max {
			xStr := strconv.Itoa(x)
			if xStr[:len(xStr)/2] == xStr[len(xStr)/2:] {
				s.InvalidIds = append(s.InvalidIds, x)
			}
			x++
		}
	}
	total := 0
	for _, i := range s.InvalidIds {
		total += i
	}
	fmt.Printf("Part 1: %d\n", total)
}

func (s *Solution) solvePartTwoBase() {
	total := 0
	for _, r := range s.Ranges {
		total += r.checkRange()
	}
	fmt.Printf("Part 2: %d\n", total)
}

func (r *Range) checkRange() int {
	if len(r.MinStr) < 2 && len(r.MaxStr) == len(r.MinStr) {
		return 0
	}
	x := r.Min
	total := 0
	for x <= r.Max {
		xStr := strconv.Itoa(x)
		for i := len(xStr) / 2; i > 0; i-- {
			splitted := strings.SplitAfter(xStr, xStr[:i])
			actualSplit := []string{}
			for _, s := range splitted {
				if s == "" {
					continue
				}
				actualSplit = append(actualSplit, s)
			}
			if compareIds(actualSplit) {
				total += x
				break
			}
		}
		x++
	}
	return total
}

func compareIds(ids []string) bool {
	v := ""
	matches := 0
	for idx, r := range ids {
		if idx == 0 {
			v = r
		} else {
			if r == v {
				matches += 1
			}
		}
	}
	return matches == len(ids)-1
}

func solvePartOne() {
	s := &Solution{}
	s.parseInput(input)
	s.solvePartOne()
}

func solvePartTwo() {
	s := &Solution{}
	s.parseInput(input)
	s.solvePartTwoBase()
}
