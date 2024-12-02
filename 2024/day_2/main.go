package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Schematic struct {
	Rows  []Row
	Total int
}

type Row struct {
	Values      []int
	ValueString string
	Safe        bool
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func (s *Schematic) ParseInput(input string) {
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, row := range splitted {
		r := Row{
			Safe:        false,
			ValueString: row,
		}
		s.Rows = append(s.Rows, r)
	}
}

func (s *Schematic) CalculateValues() {
	for i, row := range s.Rows {
		splitted := strings.Split(row.ValueString, " ")
		for _, s := range splitted {
			val, _ := strconv.Atoi(s)
			row.Values = append(row.Values, val)
		}
		s.Rows[i] = row
	}
}

func (row *Row) Validate() bool {

	increasing := false
	init := false
	for i, val := range row.Values {
		if len(row.Values) == i+1 {
			//fmt.Printf("Row safe: %v\n", row.Values)
			row.Safe = true
			break
		}
		nextVal := row.Values[i+1]
		if !init {
			increasing = val < nextVal
			init = true
		}
		localIncreasing := val < nextVal

		diff := int(math.Abs(float64(val - nextVal)))
		if (localIncreasing != increasing) || (diff < 1 || diff > 3) {
			//fmt.Printf("Row is not safe: %v, %v, %v\n", row.Values, localIncreasing != increasing, diff)
			row.Safe = false
			break
		}
	}
	return row.Safe
}

func (s *Schematic) ValidateRows() {
	for i, row := range s.Rows {
		safe := row.Validate()
		s.Rows[i].Safe = safe
	}
}

func (s *Schematic) CalculateTotal() {
	for _, row := range s.Rows {
		if row.Safe {
			s.Total += 1
		}
	}
}

func (s *Schematic) TestFaultyRows() {
	for i, row := range s.Rows {
		if row.Safe {
			continue
		}
		for i := range row.Values {
			localCopy := Row{}
			for j := range row.Values {
				if i != j {
					localCopy.Values = append(localCopy.Values, row.Values[j])
				}
			}
			safe := localCopy.Validate()
			if safe {
				row.Safe = true
				break
			}
		}
		s.Rows[i] = row
	}
}

func solvePartOne(input string) {
	s := Schematic{Total: 0}
	s.ParseInput(input)
	s.CalculateValues()
	s.ValidateRows()
	s.CalculateTotal()
	fmt.Printf("Part 1: %d\n", s.Total)
}

func solvePartTwo(input string) {
	s := Schematic{Total: 0}
	s.ParseInput(input)
	s.CalculateValues()
	s.ValidateRows()
	s.TestFaultyRows()
	s.CalculateTotal()
	fmt.Printf("Part 2: %d\n", s.Total)
}
