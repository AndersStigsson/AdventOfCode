package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Map struct {
	Points        []*Point
	Guard         Guard
	RowLength     int
	ColLength     int
	Loop          bool
	SolveObstacle int
}

type Point struct {
	Position        Position
	HasObstacle     bool
	HasGuard        bool
	Visited         bool
	Tested          int
	HasBeenTried    bool
	TestedDirection string
}

type Position struct {
	Row       int
	Col       int
	Direction string
}

type Guard struct {
	Position         Position
	OriginalPosition Position
	Direction        string
	Steps            []Position
}

func (m *Map) ParseInput(input string) {
	splitted := strings.Split(strings.TrimRight(input, "\n"), "\n")
	m.RowLength = len(splitted)
	m.ColLength = len(splitted[0])
	for i, s := range splitted {
		for j, r := range s {
			p := &Point{}
			switch r {
			case '.':
				p.HasObstacle = false
			case '#':
				p.HasObstacle = true
			case '^':
				g := Guard{}
				g.Direction = "up"
				g.Position = Position{Col: j, Row: i}
				g.OriginalPosition = Position{Col: j, Row: i}
				g.Steps = append(g.Steps, g.Position)
				m.Guard = g
				p.HasObstacle = false
				p.HasGuard = true
			}
			pos := Position{Col: j, Row: i}
			p.Position = pos
			m.Points = append(m.Points, p)
		}
	}
}

func (m *Map) GetPointAtPosition(col int, row int) (int, *Point) {
	for i, p := range m.Points {
		if p.Position.Col == col && p.Position.Row == row {
			return i, p
		}
	}
	return -1, nil
}

func (m *Map) CheckAmountOfMultipleTests() int {
	total := 0
	for _, p := range m.Points {
		if p.Tested > 1 {
			total += 1
		}
	}
	return total
}

func (g Guard) CheckIfPreviouslyVisited(pos Position) bool {
	if slices.ContainsFunc(g.Steps, func(s Position) bool {
		return ((s.Col != g.OriginalPosition.Col && s.Row != g.OriginalPosition.Row) || len(g.Steps) > 1) && s.Col == pos.Col && s.Row == pos.Row && s.Direction == pos.Direction
	}) {
		// fmt.Printf("Has already visited (%d, %d) with direction %s\n", pos.Row, pos.Col, pos.Direction)
		return true
	}
	return false
}

func (m *Map) Traverse() {
	finished := false
	// finishedX := -1
	// finishedY := -1
	for {

		g := m.Guard
		// var lastTestedObstacle *Point
		switch g.Direction {
		case "up":
			idx, p := m.GetPointAtPosition(g.Position.Col, g.Position.Row-1)
			if p == nil {
				finished = true
				// finishedX = g.Position.Col
				// finishedY = g.Position.Row - 1
				break
			}
			if p.HasObstacle {
				g.Direction = "right"
				p.Tested += 1
				p.TestedDirection = "up"
				// lastTestedObstacle = p
			} else {
				pos := g.Position
				pos.Row = pos.Row - 1
				pos.Direction = "up"
				if g.CheckIfPreviouslyVisited(pos) {
					m.Loop = true
					finished = true
					break
				}
				g.Position = pos
				step := pos
				step.Direction = g.Direction
				g.Steps = append(g.Steps, step)
				p.Visited = true
			}
			m.Points[idx] = p
		case "down":
			idx, p := m.GetPointAtPosition(g.Position.Col, g.Position.Row+1)
			if p == nil {
				finished = true
				// finishedX = g.Position.Col
				// finishedY = g.Position.Row + 1
				break
			}
			if p.HasObstacle {
				g.Direction = "left"
				p.Tested += 1
			} else {
				pos := g.Position
				pos.Row = pos.Row + 1
				pos.Direction = "down"
				if g.CheckIfPreviouslyVisited(pos) {
					m.Loop = true
					finished = true
					break
				}
				g.Position = pos
				step := pos
				step.Direction = g.Direction
				g.Steps = append(g.Steps, step)
				p.Visited = true
				p.TestedDirection = "down"
				// lastTestedObstacle = p
			}
			m.Points[idx] = p
		case "right":
			idx, p := m.GetPointAtPosition(g.Position.Col+1, g.Position.Row)
			if p == nil {
				finished = true
				// finishedX = g.Position.Col + 1
				// finishedY = g.Position.Row
				break
			}
			if p.HasObstacle {
				g.Direction = "down"
				p.Tested += 1
			} else {
				pos := g.Position
				pos.Col = pos.Col + 1
				pos.Direction = "right"
				if g.CheckIfPreviouslyVisited(pos) {
					m.Loop = true
					finished = true
					break
				}
				g.Position = pos
				step := pos
				step.Direction = g.Direction
				g.Steps = append(g.Steps, step)
				p.Visited = true
				p.TestedDirection = "right"
				// lastTestedObstacle = p
			}
			m.Points[idx] = p
		case "left":
			idx, p := m.GetPointAtPosition(g.Position.Col-1, g.Position.Row)
			if p == nil {
				finished = true
				// finishedX = g.Position.Col - 1
				// finishedY = g.Position.Row
				break
			}
			if p.HasObstacle {
				g.Direction = "up"
				p.Tested += 1
			} else {
				pos := g.Position
				pos.Col = pos.Col - 1
				pos.Direction = "left"
				if g.CheckIfPreviouslyVisited(pos) {
					m.Loop = true
					finished = true
					break
				}
				g.Position = pos
				step := pos
				step.Direction = g.Direction
				g.Steps = append(g.Steps, step)
				p.Visited = true
				p.TestedDirection = "left"
				// lastTestedObstacle = p
			}
			m.Points[idx] = p
		}
		m.Guard = g
		// m.Print()
		// fmt.Printf("Guard is: %v\n", m.Guard)
		if finished {
			// fmt.Printf("Walked out of map at pos (%d, %d) \n", finishedY, finishedX)
			break
		}
	}
}

func (m *Map) Reset() {
	for i, p := range m.Points {
		p.Tested = 0
		m.Points[i] = p
	}
}

func (m *Map) AddObstacles() {
	originalPosition := Position{Row: m.Guard.OriginalPosition.Row, Col: m.Guard.Position.Col}

	for i, p := range m.Points {
		m.Guard.Direction = "up"
		m.Guard.Position.Col = originalPosition.Col
		m.Guard.Position.Row = originalPosition.Row
		m.Guard.Steps = []Position{}
		// m.Print()
		m.Reset()
		if !p.HasGuard && !p.HasObstacle {
			fmt.Printf("Testing with obstacle at (%d, %d)\n", p.Position.Row, p.Position.Col)
			p.HasObstacle = true
			m.Points[i] = p
			// m.Print()
			m.Traverse()
			if !m.Loop {
				p.HasObstacle = false
			} else {
				m.Loop = false
				m.SolveObstacle += 1
				p.HasObstacle = false
			}
			m.Points[i] = p
		}
	}
}

func (m *Map) Solve() int {
	sum := 0
	for _, p := range m.Points {
		if p.Visited {
			sum += 1
		}
	}
	return sum
}

func (m *Map) Print() {
	for _, p := range m.Points {
		if p.Position.Col == m.Guard.Position.Col && p.Position.Row == m.Guard.Position.Row {
			if m.Guard.Direction == "up" {
				fmt.Print("^")
			}
			if m.Guard.Direction == "left" {
				fmt.Print("<")
			}
			if m.Guard.Direction == "right" {
				fmt.Print(">")
			}
			if m.Guard.Direction == "down" {
				fmt.Print("v")
			}
		} else if p.HasObstacle {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if p.Position.Col == m.ColLength-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}

func solvePartOne(input string) {
	m := Map{}
	m.ParseInput(input)
	m.Print()
	m.Traverse()
	res := m.Solve()
	fmt.Printf("Loop: %v\n", m.Loop)
	fmt.Printf("Part 1: %d\n", res)

}

func solvePartTwo(input string) {
	m := Map{}
	m.ParseInput(input)
	fmt.Printf("Loop: %v\n", m.Loop)
	m.Print()
	m.AddObstacles()
	fmt.Printf("Part 2: %d\n", m.SolveObstacle)

}

func main() {
	// solvePartOne(input)
	solvePartTwo(input)
}
