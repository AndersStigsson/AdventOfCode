package main

import (
	"fmt"
	"os"
)

var (
	currentState []byte
	index        int
	partTwo      []byte
)

func main() {
	file, err := os.ReadFile("./banana")
	if err != nil {
		panic(err.Error())
	}
	handleStringPartOne(file)
	handlePartTwo(file)
}

func handleStringPartOne(content []byte) {
	currentState = append(currentState, content[:4]...)
	index = 4
	check := checkIfUnique(currentState)
	if check {
		fmt.Printf("Solved: index %d", index)
	}

	for _, v := range content[4:] {
		currentState = currentState[1:]
		currentState = append(currentState, v)
		index += 1
		check := checkIfUnique(currentState)
		if check {
			fmt.Printf("Solved Part One: index %d\n", index)
			break
		}
	}
}

func handlePartTwo(content []byte) {
	partTwo = append(partTwo, content[:14]...)
	index = 14
	check := checkIfUnique(partTwo)
	if check {
		fmt.Printf("Solved: index %d", index)
	}

	for _, v := range content[14:] {
		partTwo = partTwo[1:]
		partTwo = append(partTwo, v)
		index += 1
		check := checkIfUnique(partTwo)
		if check {
			fmt.Printf("Solved Part Two: index %d\n", index)
			break
		}
	}
}

func checkIfUnique(state []byte) bool {
	for k, v := range state {
		for i := k + 1; i < len(state); i++ {
			if state[i] == v {
				return false
			}
		}
	}
	return true
}
