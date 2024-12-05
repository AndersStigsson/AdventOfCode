package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var target []rune = []rune{'X', 'M', 'A', 'S'}

type Matrix struct {
	Nodes []Node
	Rows  []string
	Sum   int
}
type Node struct {
	Character  rune
	Neighbours []*Node
	Row        int
	Col        int
	HasMatch   bool
	Sum        int
	IsStart    bool
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func (n *Node) CheckIfNodeHasCharacter(index int) bool {
	return n.Character == target[index]
}

func (n *Node) FindCharacterInNeighboursWithDirection(index int, direction string, m *Matrix, descending bool) bool {
	if !n.CheckIfNodeHasCharacter(index) {
		// fmt.Printf("Node (%d, %d) does not have %s\n", n.Row, n.Col, string(target[index]))
		return false
	}
	if index == 3 && !descending {
		return true
	}

	if index == 1 && descending {
		return true
	}
	var nb *Node
	switch direction {
	case "right":
		nb = m.FindNode(n.Row, n.Col+1)
	case "left":
		nb = m.FindNode(n.Row, n.Col-1)
	case "up":
		nb = m.FindNode(n.Row-1, n.Col)
	case "down":
		nb = m.FindNode(n.Row+1, n.Col)
	case "down-right":
		nb = m.FindNode(n.Row+1, n.Col+1)
	case "down-left":
		nb = m.FindNode(n.Row+1, n.Col-1)
	case "up-right":
		nb = m.FindNode(n.Row-1, n.Col+1)
	case "up-left":
		nb = m.FindNode(n.Row-1, n.Col-1)
	}
	if nb != nil {
		nextIndex := index + 1
		if descending {
			nextIndex = index - 1
		}
		return nb.FindCharacterInNeighboursWithDirection(nextIndex, direction, m, descending)
	}

	return false
}

func (n *Node) FindCharacterInNeighbours(index int, m *Matrix) bool {
	if !n.CheckIfNodeHasCharacter(index) {
		return false
	}
	if index == 3 {
		return true
	}
	sum := 0
	for _, nb := range n.Neighbours {
		direction := "right"
		if nb.Row > n.Row {
			if nb.Col == n.Col {
				direction = "down"
			} else if nb.Col > n.Col {
				direction = "down-right"
			} else {
				direction = "down-left"
			}
		} else if nb.Row < n.Row {
			if nb.Col == n.Col {
				direction = "up"
			} else if nb.Col > n.Col {
				direction = "up-right"
			} else {
				direction = "up-left"
			}
		} else {
			if nb.Col < n.Col {
				direction = "left"
			} else if nb.Col > n.Col {
				direction = "right"
			}
		}
		if nb.FindCharacterInNeighboursWithDirection(index+1, direction, m, false) {
			sum += 1
		}
	}
	if index == 0 {
		n.HasMatch = true
		n.Sum = sum
	}
	return sum > 0
}

func (m *Matrix) FindOccurences() {
	for i, n := range m.Nodes {
		if n.Character == target[0] {
			n.FindCharacterInNeighbours(0, m)
			m.Nodes[i] = n
		}
	}
}

func (n *Node) FindNeighbourByDirection(direction string) *Node {
	row := n.Row
	col := n.Col
	switch direction {
	case "up-left":
		row = row - 1
		col = col - 1
	case "up-right":
		row = row - 1
		col = col + 1
	case "down-left":
		row = row + 1
		col = col - 1
	case "down-right":
		row = row + 1
		col = col + 1
	}
	for _, nb := range n.Neighbours {
		if nb.Row == row && nb.Col == col {
			return nb
		}
	}
	return nil
}

func (m *Matrix) FindPartTwoOccurences() {
	for i, n := range m.Nodes {
		if n.Character == target[2] {
			matches := 0
			nb := n.FindNeighbourByDirection("up-left")
			var ulC rune
			if nb != nil {
				ulC = nb.Character
			} else {
				continue
			}
			nb2 := n.FindNeighbourByDirection("down-right")
			if nb2 != nil {
				if ulC == 'M' {
					if nb2.Character == 'S' {
						matches += 1
					}
				} else if ulC == 'S' {
					if nb2.Character == 'M' {
						matches += 1
					}
				}
			}

			nb = n.FindNeighbourByDirection("up-right")
			if nb != nil {
				ulC = nb.Character
			} else {
				continue
			}

			nb2 = n.FindNeighbourByDirection("down-left")
			if nb2 != nil {
				if ulC == 'M' {
					if nb2.Character == 'S' {
						matches += 1
					}
				} else if ulC == 'S' {
					if nb2.Character == 'M' {
						matches += 1
					}
				}
			}
			if matches == 2 {
				n.HasMatch = true
				m.Nodes[i] = n
			}
		}
	}
}

func (n *Node) AddNeighbours(m Matrix) {
	rs := []int{n.Row - 1, n.Row, n.Row + 1}
	cs := []int{n.Col - 1, n.Col, n.Col + 1}
	for _, row := range rs {
		for _, col := range cs {
			if row == n.Row && col == n.Col {
				continue
			}
			nb := m.FindNode(row, col)
			if nb != nil {
				n.Neighbours = append(n.Neighbours, nb)
			}

		}
	}
}

func (m *Matrix) FindNode(row int, col int) *Node {
	for _, n := range m.Nodes {
		if n.Row == row && n.Col == col {
			return &n
		}
	}
	if row == -1 || col == -1 {
		return nil
	}
	if row > len(m.Rows)-1 {
		return nil
	}
	colLen := len(m.Rows[row])
	if col > colLen-1 {
		return nil
	}

	n := Node{Character: rune(m.Rows[row][col]), Row: row, Col: col}
	m.Nodes = append(m.Nodes, n)
	return &n
}

func (m *Matrix) CreateNeighbours() {
	for i, n := range m.Nodes {
		n.AddNeighbours(*m)
		m.Nodes[i] = n
	}
}

func (m *Matrix) parseInput(input string) {
	splitted := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for i, s := range splitted {
		for j, r := range s {
			n := Node{Character: r, Row: i, Col: j}
			m.Nodes = append(m.Nodes, n)
		}
		m.Rows = append(m.Rows, s)
	}
}

func (m *Matrix) Solve(partTwo bool) {
	sum := 0
	for _, n := range m.Nodes {
		if n.HasMatch {
			if partTwo {
				sum += 1
			} else {
				sum += n.Sum
			}
		}
	}
	m.Sum = sum
}

func (m *Matrix) Print() {
	for _, n := range m.Nodes {
		// fmt.Printf("Node %d: Row: %d, Col: %d, Neighbours %d\n", i, n.Row, n.Col, len(n.Neighbours))
		if n.HasMatch {
			fmt.Printf("X")
		} else {
			fmt.Printf(".")
		}

		if n.Col == len(m.Rows[n.Row])-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func (m *Matrix) PrintNeighbours() {
	for _, n := range m.Nodes {
		fmt.Printf("Node: (%d, %d) %s  Neighbours:", n.Row, n.Col, string(n.Character))
		for _, nb := range n.Neighbours {
			fmt.Printf("(%d, %d) %s, ", nb.Row, nb.Col, string(nb.Character))
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) PrintOriginal() {
	// result := map[int]map[int]string{}
	// for i := range m.Rows {
	// 	result[i] = make(map[int]string, len(m.Rows[i]))
	// }

	for _, n := range m.Nodes {
		fmt.Printf("%s", string(n.Character))
		if n.Col == len(m.Rows[n.Row])-1 {
			fmt.Printf("\n")
		}
	}

	// for i := range result {
	// 	for j := range result[i] {
	// 		fmt.Printf("%s", result[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }
}

func solvePartOne(input string) {
	m := Matrix{}
	m.parseInput(input)
	m.CreateNeighbours()
	m.FindOccurences()
	// m.Print()
	// m.PrintNeighbours()
	//m.PrintOriginal()
	m.Solve(false)
	fmt.Printf("Part 1: %d\n", m.Sum)
}

func solvePartTwo(input string) {
	m := Matrix{}
	m.parseInput(input)
	m.CreateNeighbours()
	m.FindPartTwoOccurences()
	m.Solve(true)
	fmt.Printf("Part 2: %d\n", m.Sum)
}
