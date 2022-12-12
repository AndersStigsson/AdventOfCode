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

type Position struct {
	X int
	Y int
}

type Instruction struct {
	Direction string
}

type InstructionSet struct {
	Instructions []Instruction
}

type Point struct {
	Visited bool
	Pos     Position
}

type Knot struct {
	Point Point
}

type State struct {
	Head          Knot
	Tail          Knot
	Knots         []Knot
	VisitedPoints map[Position]bool
}

func (T *Knot) distanceToHead(head Knot) bool {
	// if Abs(head.Point.Pos.X - T.Point.Pos.X) > 1 ||
	//     Abs(head.Point.Pos.Y - T.Point.Pos.Y) > 1
	return int(math.Abs(float64(head.Point.Pos.X-T.Point.Pos.X))) > 1 ||
		int(math.Abs(float64(head.Point.Pos.Y-T.Point.Pos.Y))) > 1
}

func (State *State) doStuff(instructions InstructionSet) {
	for _, inst := range instructions.Instructions {
		State.executeInstruction(inst)
	}
}

func (State *State) moveTailIfNeeded(direction string) {
	if State.Tail.distanceToHead(State.Head) {
		State.moveTail(State.Head)
	}
}

func (State *State) moveAllKnotsIfNeeded() {
	startKnot := State.Head
	earlyStop := false
	for i, v := range State.Knots {
		if !v.distanceToHead(startKnot) {
			earlyStop = true
			break
		}
		State.moveKnot(i, startKnot)
		startKnot = State.Knots[i]
	}

	startKnot = State.Knots[len(State.Knots)-1]
	if !earlyStop && State.Tail.distanceToHead(startKnot) {
		State.moveTail(startKnot)
	}
}

func (State *State) moveKnot(index int, head Knot) {
	tail := State.Knots[index]
	if head.Point.Pos.Y > tail.Point.Pos.Y {
		tail.Point.Pos.Y++
	} else if head.Point.Pos.Y < tail.Point.Pos.Y {
		tail.Point.Pos.Y--
	}
	if head.Point.Pos.X > tail.Point.Pos.X {
		tail.Point.Pos.X++
	} else if head.Point.Pos.X < tail.Point.Pos.X {
		tail.Point.Pos.X--
	}
	State.Knots[index] = tail
}

func (State *State) moveTail(head Knot) {
	if head.Point.Pos.Y > State.Tail.Point.Pos.Y {
		State.Tail.Point.Pos.Y++
	} else if head.Point.Pos.Y < State.Tail.Point.Pos.Y {
		State.Tail.Point.Pos.Y--
	}
	if head.Point.Pos.X > State.Tail.Point.Pos.X {
		State.Tail.Point.Pos.X++
	} else if head.Point.Pos.X < State.Tail.Point.Pos.X {
		State.Tail.Point.Pos.X--
	}

	State.VisitedPoints[State.Tail.Point.Pos] = true
}

func (Head *Knot) moveHead(direction string) {
	switch direction {
	case "R":
		Head.Point.Pos.X++
	case "L":
		Head.Point.Pos.X--
	case "U":
		Head.Point.Pos.Y--
	case "D":
		Head.Point.Pos.Y++
	}
}

func (State *State) executeInstruction(instruction Instruction) {
	State.Head.moveHead(instruction.Direction)
	State.moveAllKnotsIfNeeded()
	// State.moveTailIfNeeded(instruction.Direction)
}

func createInstruction(direction string, length int) []Instruction {
	instructions := []Instruction{}
	for len(instructions) < length {
		instructions = append(instructions, Instruction{Direction: direction})
	}
	return instructions
}

func parseInput() InstructionSet {
	var instructions InstructionSet
	for _, str := range strings.Split(input, "\n") {
		if len(str) == 0 {
			continue
		}
		split := strings.Split(str, " ")
		direction := split[0]
		length, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		instructions.Instructions = append(instructions.Instructions, createInstruction(direction, length)...)
	}
	return instructions
}

func main() {
	var head Knot
	var tail Knot
	knots := make([]Knot, 8)
	State := State{Head: head, Tail: tail, VisitedPoints: map[Position]bool{}, Knots: knots}
	startPoint := Point{Pos: Position{X: 0, Y: 0}}
	State.VisitedPoints[startPoint.Pos] = true
	instructions := parseInput()
	State.doStuff(instructions)
	fmt.Printf("Amount of visited points: %d\n", len(State.VisitedPoints))
}
