package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Map struct {
	FollowKey           []string
	Head                *Node
	Tail                *Node
	Nodes               map[string]*Node
	StartingNodes       []*Node
	EndingNodes         []*Node
	NumberOnEndingNodes int
}
type Node struct {
	Key         string
	LeftString  string
	RightString string
	Left        *Node
	Right       *Node
}

func main() {
	fmt.Printf("Answer part 1: %d\n", solvePartOne())
	fmt.Printf("Answer part 2: %d\n", solvePartTwo())
}

func solvePartOne() int {
	m := parseInput(input)
	m.AssignNodesToNodes()
	solutions := m.FindSolution()
	return solutions
}

func solvePartTwo() int {
	m := parseInput(input)
	m.AssignNodesToNodes()
	solutions := m.FindSolutionPartTwo()
	return solutions
}

func (m *Map) FindSolutionPartTwo() int {
	solveIndices := []int{}
	for _, v := range m.StartingNodes {
		solveIndices = append(solveIndices, v.FindEndingLine(m.FollowKey))
	}

	return findCommonMultiple(solveIndices)
}

func (n *Node) FindEndingLine(followKey []string) int {
	index := 0
	nc := n
	for {
		rl := followKey[index%len(followKey)]
		if rl == "R" {
			// fmt.Printf("Node %s moves to %s\n", n.Key, n.Right.Key)
			nc = nc.Right
		} else if rl == "L" {
			// fmt.Printf("Node %s moves to %s\n", n.Key, n.Left.Key)
			nc = nc.Left
		}
		index += 1
		if string(nc.Key[2]) == "Z" {
			break
		}
	}
	return index
}

func (m *Map) FindSolution() int {
	index := 0
	n := m.Head
	for {
		rl := m.FollowKey[index%len(m.FollowKey)]
		if rl == "R" {
			n = n.Right
		} else if rl == "L" {
			n = n.Left
		}
		index += 1
		if n.Key == "ZZZ" {
			break
		}
	}

	return index
}

func (m *Map) AssignNodesToNodes() {
	for i, n := range m.Nodes {
		n.Left = m.Nodes[n.LeftString]
		n.Right = m.Nodes[n.RightString]
		m.Nodes[i] = n
	}
}

func findCommonMultiple(values []int) int {
	return LCM(values[0], values[1], values[2:]...)
}

func getFollowKey(row string) []string {
	retval := []string{}
	for _, v := range row {
		retval = append(retval, string(v))
	}

	return retval
}

func parseInput(input string) Map {
	m := Map{}
	m.Nodes = map[string]*Node{}

	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		if v == "" {
			continue
		}
		if !strings.Contains(v, "=") {
			str := strings.TrimSpace(v)
			m.FollowKey = append(m.FollowKey, getFollowKey(str)...)
			continue
		}
		n := &Node{}
		split := strings.Split(v, " = ")
		n.Key = split[0]
		if n.Key == "AAA" {
			m.Head = n
		} else if n.Key == "ZZZ" {
			m.Tail = n
		}
		if string(n.Key[2]) == "A" {
			m.StartingNodes = append(m.StartingNodes, n)
		} else if string(n.Key[2]) == "Z" {
			m.EndingNodes = append(m.EndingNodes, n)
		}
		targets := strings.Split(split[1], ",")
		n.LeftString = strings.Trim(targets[0], "() ")
		n.RightString = strings.Trim(targets[1], "() ")
		m.Nodes[n.Key] = n
	}

	return m
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
