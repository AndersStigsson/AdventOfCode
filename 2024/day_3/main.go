package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Calculation struct {
	Left  int
	Right int
}

type Memory struct {
	Calculations []Calculation
	Rows         []Row
	Sum          int
}

type Row struct {
	Text    string
	SubRows []SubRow
}

type SubRow struct {
	Text string
	Do   bool
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func parseInput(input string) *Memory {
	m := Memory{}
	splitted := strings.Split(strings.TrimRight(input, "\n"), "\n")
	rows := []Row{}
	for _, s := range splitted {
		r := Row{Text: s}
		rows = append(rows, r)
	}
	m.Rows = rows
	return &m
}

func (r *Row) FindMulPartTwo() {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	split := re.FindAllIndex([]byte(r.Text), -1)
	lastIdx := -1
	lastDo := false
	for i, idx := range split {
		sub := SubRow{}
		if i == 0 {
			substr := string(r.Text)[:idx[1]]
			sub.Text = substr
			if strings.Contains(substr, "don't()") {
				sub.Do = false
				lastDo = false
			} else {
				sub.Do = true
			}
		} else {
			substr := string(r.Text)[lastIdx:idx[1]]
			sub.Text = substr
			if strings.Contains(substr, "don't()") {
				sub.Do = false
				lastDo = false
			} else if strings.Contains(substr, "do()") {
				sub.Do = true
				lastDo = true
			} else if lastDo {
				sub.Do = true
			}
		}
		r.SubRows = append(r.SubRows, sub)
		lastIdx = idx[1]
	}
}

func (m *Memory) FindMulPartTwo() {
	for i, r := range m.Rows {
		r.FindMulPartTwo()
		m.Rows[i] = r
	}
}

func (r *Row) FindMul() [][]int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAll([]byte(r.Text), -1)
	re = regexp.MustCompile(`\d+`)
	mulMatches := [][]int{}
	for _, match := range matches {
		localmm := []int{}
		mm := re.FindAll(match, -1)
		for _, m := range mm {
			val, err := strconv.Atoi(string(m))
			if err != nil {
				panic(err)
			}
			localmm = append(localmm, val)
		}
		mulMatches = append(mulMatches, localmm)
	}
	return mulMatches
}

func (m *Memory) FindMul() {
	calculations := []Calculation{}
	for _, row := range m.Rows {
		mulByRow := row.FindMul()
		for _, mbr := range mulByRow {
			calc := Calculation{Left: mbr[0], Right: mbr[1]}
			calculations = append(calculations, calc)
		}
	}
	m.Calculations = calculations
}

func (c *Calculation) Solve() int {
	return c.Left * c.Right
}

func (m *Memory) Solve() {
	sum := 0
	for _, calc := range m.Calculations {
		sum += calc.Solve()
	}
	m.Sum = sum
}

func (s *SubRow) GetCalculations() Calculation {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	match := re.Find([]byte(s.Text))
	re = regexp.MustCompile(`\d+`)
	mm := re.FindAll(match, -1)
	c := Calculation{}
	for i, m := range mm {
		val, err := strconv.Atoi(string(m))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			c.Left = val
		} else {
			c.Right = val
		}
	}
	return c
}

func (m *Memory) CalculateSubRows() {
	for _, row := range m.Rows {
		for _, sr := range row.SubRows {
			if sr.Do {
				calc := sr.GetCalculations()
				m.Calculations = append(m.Calculations, calc)
			}
		}
	}
}

func solvePartOne(input string) {
	m := parseInput(input)
	m.FindMul()
	m.Solve()
	fmt.Printf("Part 1: %d\n", m.Sum)
}

func solvePartTwo(input string) {
	m := parseInput(input)
	m.FindMulPartTwo()
	m.CalculateSubRows()
	m.Solve()
	fmt.Printf("Part 2: %d\n", m.Sum)
}
