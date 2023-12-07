package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
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
	now := time.Now()
	fmt.Printf("Part 1: %d, time: %vms\n", solvePartOne(), float64(time.Since(now).Microseconds())/float64(1000))
	now = time.Now()
	fmt.Printf("Part 2: %d, time: %vms\n", solvePartTwo(), float64(time.Since(now))/float64(time.Millisecond))
}

func solvePartOne() int {
	c := parseInput(input)
	solutions := []int{}
	for _, r := range c.Races {
		solutions = append(solutions, r.FindSolutions())
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
	ts := ""
	ds := ""
	for _, r := range c.Races {
		ts += fmt.Sprintf("%d", r.Time)
		ds += fmt.Sprintf("%d", r.Distance)
	}

	t, _ := strconv.Atoi(ts)
	d, _ := strconv.Atoi(ds)

	r := Race{Time: t, Distance: d}
	solution := r.FindSolutions()

	return solution
}

func (r *Race) FindSolutions() int {
	solutions := 0
	p := findPivotValue(r.Time, r.Distance)
	solutions = (r.Time - 2*p) + 1
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

// v is the value to test
// t is the time
// d is the distance
func solvesFn(v, t, d int) bool {
	return v*(t-v) > d
}

// t being the time for the race, d being the distance
func findPivotValue(t int, d int) int {
	pivot := 0
	for i := 1; i < int(t/2); i++ {
		test := float64(t) / math.Pow(2, float64(i))
		if !solvesFn(int(test), t, d) {
			pivot = int(test)
			break
		}
	}

	for i := pivot; i < int(t/2); i++ {
		if solvesFn(i, t, d) {
			pivot = i
			break
		}
	}
	return pivot
}

func rangeFn(start int, stop int) []int {
	ret := []int{}
	for i := start; i < stop; i++ {
		ret = append(ret, i)
	}
	return ret
}
