package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed banana.txt
var input string

type Instruction struct {
	Operation        string
	Value            int
	CyclesToComplete int
}

type InstructionSet struct {
	Instructions []Instruction
}

type State struct {
	Cycle            int
	RegisterX        int
	CycleToCheck     int
	CyclesOfInterest []int
	Pixels           [][]string
}

func createInstruction(split []string) Instruction {
	operation := split[0]
	additionValue := 0
	cyclesToComplete := 1

	if operation == "addx" {
		cyclesToComplete = 2
		value, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		additionValue = value
	}
	return Instruction{Operation: operation, Value: additionValue, CyclesToComplete: cyclesToComplete}
}

func parseInput() InstructionSet {
	var instructions InstructionSet
	for _, str := range strings.Split(input, "\n") {
		if len(str) == 0 {
			continue
		}
		split := strings.Split(str, " ")
		instructions.Instructions = append(instructions.Instructions, createInstruction(split))
	}
	return instructions
}

func (State *State) AddToCycleSlice() {
	fmt.Printf("At Cycle %d, RegisterX %d\n", State.Cycle, State.RegisterX)
	State.CyclesOfInterest = append(State.CyclesOfInterest, State.Cycle*State.RegisterX)
}

func (State *State) CheckIfCycleMatches() bool {
	return State.Cycle == State.CycleToCheck
}

func (State *State) HandleSprite() {
	row := int(math.Floor(float64(State.Cycle) / 40))
	column := State.Cycle % 40
	difference := State.RegisterX - column
	fmt.Printf("[%d, %d]\n", row, column)
	if row > 5 {
		return
	}
	fmt.Printf("column = %d, State.RegisterX = %d, difference: %d\n", column, State.RegisterX, difference)
	if column >= State.RegisterX-1 && column <= State.RegisterX+1 {
		State.Pixels[row][column] = "#"
	} else {
		State.Pixels[row][column] = "."
	}
}

func (State *State) ExecuteInstruction(instruction Instruction) {
	for i := 0; i < instruction.CyclesToComplete; i++ {
		State.HandleSprite()
		State.Cycle++
		if State.CheckIfCycleMatches() {
			State.AddToCycleSlice()
			State.CycleToCheck += 40
		}
	}

	State.RegisterX += instruction.Value
}

func (State *State) CalculateCycleValues() int {
	value := 0
	for _, v := range State.CyclesOfInterest {
		// fmt.Printf("Cycle %d has value: %d\n", i, v)
		value += v
	}
	return value
}

func (State *State) printPixels() {
	for i := range State.Pixels {
		for _, v := range State.Pixels[i] {
			fmt.Print(v)
		}
		fmt.Printf("\n")
	}
}

func (State *State) doStuff(instructions []Instruction) {
	for _, inst := range instructions {
		State.ExecuteInstruction(inst)
	}
}

func main() {
	pixels := make([][]string, 6)
	for i := range pixels {
		pixels[i] = make([]string, 40)
	}
	state := State{Cycle: 0, RegisterX: 1, CycleToCheck: 20, Pixels: pixels}
	instructions := parseInput()
	state.doStuff(instructions.Instructions)
	// value := state.CalculateCycleValues()
	// fmt.Printf("Value is: %d", value)
	state.printPixels()
}
