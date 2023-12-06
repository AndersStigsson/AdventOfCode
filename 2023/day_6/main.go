package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Competition struct {
	Races []Race
}

type Race struct {
	Time     int
	Distance int
}

func main() {
	fmt.Printf("Part 1: %d\n", solvePartOne())
	fmt.Printf("Part 2: %d\n", solvePartTwo())
	// fmt.Printf("Solutions: %d\n", solveForX(9, 7))
	// fmt.Printf("Solutions: %d\n", solveForX(40, 15))
	// fmt.Printf("Solutions: %d\n", solveForX(200, 30))
}

func solvePartOne() int {
	c := parseInput(input)
	solutions := []int{}
	for _, r := range c.Races {
		solutions = append(solutions, r.SolveForX())
	}

	if len(solutions) == 0 {
		return 0
	}

	tot := 1
	for _, v := range solutions {
		tot *= v
	}

	return tot
}

func solvePartTwo() int {
	c := parseInput(input)
	fmt.Printf("Races %v\n", c.Races)
	ts := ""
	ds := ""
	for _, r := range c.Races {
		ts += fmt.Sprintf("%d", r.Time)
		ds += fmt.Sprintf("%d", r.Distance)
	}

	t, _ := strconv.Atoi(ts)
	d, _ := strconv.Atoi(ds)

	r := Race{Time: t, Distance: d}
	solution := r.SolveForX()

	return solution
}

func (r *Race) SolveForX() int {
	solutions := 0
	for _, v := range rangeFn(0, r.Time+1) {
		if v*(r.Time-v) > r.Distance {
			solutions++
		}
	}
	return solutions
}

func parseInput(input string) Competition {
	c := Competition{}
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, str := range splitted {
		split := strings.Split(strings.Trim(str, " "), ":")
		vals := split[1]
		split = strings.Split(strings.TrimSpace(vals), " ")
		idx := 0
		for _, s := range split {
			if s == "" {
				continue
			}
			t, _ := strconv.Atoi(s)
			if len(c.Races) > idx {
				c.Races[idx].Distance = t
			} else {
				r := Race{Time: t}
				c.Races = append(c.Races, r)
			}
			idx++
		}
	}

	return c
}

func solveForX(distance int, t int) int {
	solutions := 0
	for _, v := range rangeFn(0, t+1) {
		if v*(t-v) > distance {
			solutions++
		}
	}
	return solutions
}

func rangeFn(start int, stop int) []int {
	ret := []int{}
	for i := start; i < stop; i++ {
		ret = append(ret, i)
	}
	return ret
}
