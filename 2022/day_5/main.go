package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State [][]string

func (input State) parseInputLine(line []byte) {
	idx := 0
	for k, v := range line {
		if v >= 'A' && v <= 'Z' {
			if len(input) < idx {
				input[idx] = []string{}
			}
			input[idx] = append(input[idx], string(v))
		}
		if k != 0 && k%4 == 0 {
			idx += 1
		}
	}
}

func parseFirstPart(fileContent []byte) (State, int) {
	var inputMatrix State
	initiated := false

	for k, v := range strings.Split(string(fileContent), "\n") {
		byteText := []byte(v)
		if len(byteText) == 0 {
			return inputMatrix, k
		}
		if !initiated {
			// TODO: Fix how to calculate the length based on input
			inputMatrix = make(State, 9)
			initiated = true
		}
		inputMatrix.parseInputLine(byteText)
	}
	return nil, 0
}

func main() {
	file, err := os.ReadFile("./fullfile.txt")
	if err != nil {
		panic(err.Error())
	}
	state, row := parseFirstPart(file)
	initialstate := append([][]string{}, state...)
	partOne(state, file, row)
	partTwo(initialstate, file, row)
}

func partOne(state State, fileContent []byte, idx int) {
	for k, v := range strings.Split(string(fileContent), "\n") {
		if k <= idx {
			continue
		}
		if v == "" {
			break
		}
		stuffToDo := parseSteps(v)
		state.doStuff(stuffToDo)
	}
	state.getFirstOfEachCrate("Part 1")
}

func partTwo(state State, fileContent []byte, idx int) {
	for k, v := range strings.Split(string(fileContent), "\n") {
		if k <= idx {
			continue
		}
		if v == "" {
			break
		}
		stuffToDo := parseSteps(v)
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
	prependData := append([]string(nil), initialstate[from-1][:amount]...)
	initialstate[to-1] = append(prependData, initialstate[to-1]...)
	initialstate[from-1] = initialstate[from-1][amount:]
}

func (initialstate State) getFirstOfEachCrate(part string) {
	outputString := ""
	for _, v := range initialstate {
		outputString += v[0]
	}
	fmt.Printf("%s: %s\n", part, outputString)
}
