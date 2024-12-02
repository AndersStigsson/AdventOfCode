package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func parseInput(input string) ([]int, []int) {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	left := []int{}
	right := []int{}
	for _, r := range rows {
		cols := strings.Split(r, " ")
		idx := 0
		for _, c := range cols {
			if c != "" {
				val, _ := strconv.Atoi(c)
				if idx == 0 {
					left = append(left, val)
				} else {
					right = append(right, val)
				}
				idx += 1
			}
		}
	}

	slices.SortFunc(left, func(i int, j int) int {
		return i - j
	})

	slices.SortFunc(right, func(i int, j int) int {
		return i - j
	})

	return left, right
}

func solvePartOne(input string) {
	left, right := parseInput(input)
	sum := 0
	for i := range left {
		// fmt.Printf("left is %d, right is %d\n", left[i], right[i])
		sum += calculateDistance(left[i], right[i])
	}

	fmt.Printf("ANSWER IS %d\n", sum)
}

func calculateDistance(left int, right int) int {
	if left > right {
		return left - right
	}
	if right > left {
		return right - left
	}
	return 0
}

func solvePartTwo(input string) {
	left, right := parseInput(input)
	calculationSchema := map[int]int{}
	sum := 0
	for _, i := range left {
		if _, ok := calculationSchema[i]; !ok {
			calculationSchema[i] = 0
		}
		for _, j := range right {
			if j > i {
				break
			}
			if j < i {
				continue
			}
			calculationSchema[i] += 1
		}
	}

	for l, r := range calculationSchema {
		sum += l * r
	}

	fmt.Printf("ANSWER PART 2 IS %d\n", sum)
}
