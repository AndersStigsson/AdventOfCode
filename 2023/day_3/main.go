package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type Schematic struct {
	Rows []Row
	Sum  int
}

type Row struct {
	Text          string
	Values        []Value
	SymbolIndexes []int
}

type Value struct {
	Value    int
	MaxIndex int
	MinIndex int
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func solvePartOne(input string) {
	s := Schematic{}
	s.ParseInput(input)
	s.FindSymbolsOnRows()
	s.FindValuesOnRows()
	fmt.Printf("Answer part 1: %d\n", s.SolvePartOne())
}

func solvePartTwo(input string) {
	s := Schematic{}
	s.ParseInput(input)
	s.FindSymbolsOnRows()
	s.FindValuesOnRows()
	fmt.Printf("Answer part 2: %d\n", s.SolvePartTwo())
}

func (s *Schematic) ParseInput(input string) {
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, r := range splitted {
		row := Row{Text: r}
		s.Rows = append(s.Rows, row)
	}
}

func (s *Schematic) FindValuesOnRows() {
	for idx, r := range s.Rows {
		r.FindValuesInRow()
		s.Rows[idx] = r
	}
}

func (s *Schematic) FindSymbolsOnRows() {
	for idx, r := range s.Rows {
		r.FindSymbolsInRow()
		s.Rows[idx] = r
	}
}

func (s *Schematic) SolvePartOne() int {
	sum := 0
	for idx, r := range s.Rows {
		if len(r.Values) > 0 {
			for _, v := range r.Values {
				hasAdjacentSymbol := s.CheckIfSymbolAdjacent(idx, v.MaxIndex, v.MinIndex)
				if hasAdjacentSymbol {
					fmt.Printf("Row %d should have adjacent on value %d with minIndex %d and maxIndex %d\n", idx, v.Value, v.MinIndex, v.MaxIndex)
					sum += v.Value
				}
			}
		}
	}
	return sum
}

func (s *Schematic) SolvePartTwo() int {
	sum := 0
	for idx, r := range s.Rows {
		if len(r.SymbolIndexes) > 0 {
			for _, v := range r.SymbolIndexes {
				adjacentPartnumbers := s.CheckIfAdjacentPartNumber(idx, v)
				sum += adjacentPartnumbers
			}
		}
	}
	return sum
}

func (s *Schematic) CheckIfAdjacentPartNumber(rowIdx int, symbolIdx int) int {
	power := 1

	lowest := rowIdx - 1
	highest := rowIdx + 1
	found := []int{}

	for i := lowest; i <= highest; i++ {
		row := s.Rows[i]
		for _, v := range row.Values {
			distanceMax := math.Abs(float64(v.MaxIndex - symbolIdx))
			distanceMin := math.Abs(float64(v.MinIndex - symbolIdx))

			if distanceMax <= 1 {
				found = append(found, v.Value)
				continue
			}
			if distanceMin <= 1 {
				found = append(found, v.Value)
			}
		}
	}
	if len(found) < 2 {
		return 0
	}
	for _, v := range found {
		power *= v
	}
	return power
}

func (s *Schematic) CheckIfSymbolAdjacent(idx int, maxIndex int, minIndex int) bool {
	lowest := idx - 1
	highest := idx + 1
	textIndexMin := minIndex - 1
	textIndexMax := maxIndex + 1

	if idx == 0 {
		lowest = idx
	}

	if idx == len(s.Rows)-1 {
		highest = idx
	}

	if minIndex == 0 {
		textIndexMin = minIndex
	}

	if maxIndex == len(s.Rows[highest].Text)-1 {
		textIndexMax = maxIndex
	}

	for i := lowest; i <= highest; i++ {
		row := s.Rows[i]
		exists := row.CheckIfSymbolOnIndices(textIndexMax, textIndexMin)
		if exists {
			return true
		}
	}

	return false
}

func (r *Row) CheckIfSymbolOnIndices(maxIndex int, minIndex int) bool {
	for _, s := range r.SymbolIndexes {
		if s <= maxIndex && s >= minIndex {
			return true
		}
	}
	return false
}

func (r *Row) FindSymbolsInRow() {
	for idx, t := range r.Text {
		if checkIfSymbol(t) {
			r.SymbolIndexes = append(r.SymbolIndexes, idx)
		}
	}
}

func (r *Row) FindValuesInRow() {
	minIndex := 9999999
	maxIndex := -1
	digits := ""
	for idx, t := range r.Text {
		if unicode.IsDigit(t) {
			if idx < minIndex {
				minIndex = idx
			}
			if idx > maxIndex {
				maxIndex = idx
			}
			digits += string(t)
		}
		if minIndex < 9999999 {
			if t == '.' {
				val, _ := strconv.Atoi(digits)
				v := Value{MinIndex: minIndex, MaxIndex: maxIndex, Value: val}
				r.Values = append(r.Values, v)
				minIndex = 9999999
				maxIndex = -1
				digits = ""
			} else if idx == len(r.Text)-1 {
				val, _ := strconv.Atoi(digits)
				v := Value{MinIndex: minIndex, MaxIndex: maxIndex, Value: val}
				r.Values = append(r.Values, v)
				minIndex = 9999999
				maxIndex = -1
				digits = ""
			} else if checkIfSymbol(t) {
				val, _ := strconv.Atoi(digits)
				v := Value{MinIndex: minIndex, MaxIndex: maxIndex, Value: val}
				r.Values = append(r.Values, v)
				minIndex = 9999999
				maxIndex = -1
				digits = ""
			}
		}
	}
}

func checkIfSymbol(t rune) bool {
	return !unicode.IsLetter(t) && t != '.' && !unicode.IsDigit(t)
}
