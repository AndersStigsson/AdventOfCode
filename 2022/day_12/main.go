package main

import (
	_ "embed"
	"strings"
)

//go:embed banana.txt
var input string

type Vertex struct {
	Cost    int
	Visited bool
}

type Node struct {
	X int
	Y int
}

type Graph struct {
	Vertices []Vertex
	Start    Node
	End      Node
}

type State struct {
	Graph Graph
}

func (State *State) parseInput() {
	strings.Split(input, "\n")
}

func partOne() {
	state := State{}
	state.parseInput()
}

func main() {
	partOne()
}
