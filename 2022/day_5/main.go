package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type State [][]string

func main() {
	file, _ := os.Open("./banana")
	state := State{
		{
			"D",
			"T",
			"W",
			"N",
			"L",
		},
		{
			"H",
			"P",
			"C",
		},
		{
			"J",
			"M",
			"G",
			"D",
			"N",
			"H",
			"P",
			"W",
		},
		{
			"L",
			"Q",
			"T",
			"N",
			"S",
			"W",
			"C",
		},
		{
			"N",
			"C",
			"H",
			"P",
		},
		{
			"B",
			"Q",
			"W",
			"M",
			"D",
			"N",
			"H",
			"T",
		},
		{
			"L",
			"S",
			"G",
			"J",
			"R",
			"B",
			"M",
		},
		{
			"T",
			"R",
			"B",
			"V",
			"G",
			"W",
			"N",
			"Z",
		},
		{
			"L",
			"P",
			"N",
			"D",
			"G",
			"W",
		},
	}
	// partOne(state, file)
	fmt.Println(state)
	partTwo(state, file)
}

func partOne(state State, file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		stuffToDo := parseSteps(text)
		state.doStuff(stuffToDo)
	}
	state.getFirstOfEachCrate("Part 1")
}

func partTwo(state State, file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		stuffToDo := parseSteps(text)
		state.doStuffPartTwo(stuffToDo)
	}
	state.getFirstOfEachCrate("Part 2")
}

func parseSteps(text string) []int {
	r := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	match := r.FindStringSubmatch(text)
	m1, _ := strconv.Atoi(match[1])
	m2, _ := strconv.Atoi(match[2])
	m3, _ := strconv.Atoi(match[3])
	return []int{m1, m2, m3}
}

func (initialstate State) doStuff(steps []int) {
	amount := steps[0]
	from := steps[1]
	to := steps[2]
	for i := 0; i < amount; i++ {
		initialstate[to-1] = append([]string{initialstate[from-1][i]}, initialstate[to-1]...)
	}
	initialstate[from-1] = initialstate[from-1][amount:]
}

func (initialstate State) doStuffPartTwo(steps []int) {
	amount, from, to := steps[0], steps[1], steps[2]
	fmt.Printf("Move from %d to %d amount %d, initialstate-from %v, initialstate-to %v\n", from, to, amount, initialstate[from-1], initialstate[to-1])
	prependData := append([]string(nil), initialstate[from-1][:amount]...)
	fmt.Println(prependData)
	initialstate[to-1] = append(prependData, initialstate[to-1]...)
	initialstate[from-1] = initialstate[from-1][amount:]
	fmt.Printf("Resulting: from %v, to %v\n", initialstate[from-1], initialstate[to-1])
}

func (initialstate State) getFirstOfEachCrate(part string) {
	outputString := ""
	for _, v := range initialstate {
		outputString += v[0]
	}
	fmt.Printf("%s: %s\n", part, outputString)
}
