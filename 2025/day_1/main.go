package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Rotation struct {
	Direction      string
	Value          int
	ActualMovement int
}

type Lock struct {
	Rotations    []Rotation
	NumZero      int
	CurrentValue int
	Min          int
	Max          int
	ModVal       int
}

func main() {
	l := &Lock{CurrentValue: 50, Max: 99, Min: 0, ModVal: 100}
	l.parseInput(input)

	l.solvePartOne()

	l2 := &Lock{CurrentValue: 50, Max: 99, Min: 0, ModVal: 100}
	l2.parseInput(input)

	l2.solvePartTwo()
}

func (l *Lock) parseInput(input string) {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, r := range rows {
		rot := Rotation{}
		matches := re.FindAllStringSubmatch(r, 1)
		if len(matches[0]) > 1 {
			rot.Direction = matches[0][1]
			val, err := strconv.Atoi(matches[0][2])
			if err != nil {
				panic(err)
			}
			rot.ActualMovement = val
			if rot.Direction == "L" {
				rot.ActualMovement = -1 * val
			}
			rot.Value = val
			l.Rotations = append(l.Rotations, rot)
		}

	}
}

func (l *Lock) solvePartOne() {
	for _, r := range l.Rotations {
		l.CurrentValue = (l.CurrentValue + l.ModVal + r.ActualMovement) % l.ModVal
		if l.CurrentValue == 0 {
			l.NumZero++
		}
	}

	fmt.Printf("Part 1: %d\n", l.NumZero)
}

func (l *Lock) solvePartTwo() {
	for _, r := range l.Rotations {
		prevCurrent := l.CurrentValue

		l.CurrentValue += r.ActualMovement

		l.NumZero += int(math.Floor(math.Abs(float64(l.CurrentValue) / 100)))

		if prevCurrent*l.CurrentValue < 0 {
			l.NumZero++
		}

		if l.CurrentValue == 0 {
			l.NumZero++
		}

		l.CurrentValue = l.CurrentValue % l.ModVal

	}

	fmt.Printf("Part 2: %d\n", l.NumZero)
}
