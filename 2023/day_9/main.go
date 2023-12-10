package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Oasis struct {
	History []History
}

type History struct {
	Rows []HistoryRow
}

type HistoryRow struct {
	Row []int
}

func main() {
	fmt.Printf("Answer part 1: %d\n", solvePartOne())
	fmt.Printf("Answer part 2: %d\n", solvePartTwo())
}

func solvePartOne() int {
	o := parseInput(input)
	o.CreateHistoryRows()
	o.PredictValue(false)
	// o.PrintRows()
	return o.FindSumOfLastValues(false)
}

func solvePartTwo() int {
	o := parseInput(input)
	o.CreateHistoryRows()
	o.PredictValue(true)
	// o.PrintRows()
	return o.FindSumOfLastValues(true)
}

func (o *Oasis) PrintRows() {
	for i, v := range o.History {
		fmt.Printf("History %d\n", i)
		for _, h := range v.Rows {
			fmt.Printf("%v\n", h.Row)
		}
	}
}

func (o *Oasis) FindSumOfLastValues(predictFirst bool) int {
	sum := 0
	for _, v := range o.History {
		index := len(v.Rows[0].Row) - 1
		if predictFirst {
			index = 0
		}
		sum += v.Rows[0].Row[index]
	}
	return sum
}

func (o *Oasis) PredictValue(predictFirst bool) {
	for i, h := range o.History {
		h.PredictValue(predictFirst)
		o.History[i] = h

	}
}

func (h *History) PredictValue(predictFirst bool) {
	// rows := h.Rows
	lastValue := 0
	for i := len(h.Rows) - 1; i >= 0; i-- {
		index := len(h.Rows[i].Row) - 1
		if predictFirst {
			index = 0
		}
		myValue := h.Rows[i].Row[index]
		if i+1 == len(h.Rows) {
			lastValue = 0
		} else {
			lastValue = h.Rows[i+1].Row[index]
		}
		result := myValue + lastValue
		if predictFirst {
			result = myValue - lastValue
			h.Rows[i].Row = append([]int{result}, h.Rows[i].Row...)
		} else {
			h.Rows[i].Row = append(h.Rows[i].Row, result)
		}
	}
}

func (o *Oasis) CreateHistoryRows() {
	for i, h := range o.History {
		h.CreateRows()
		o.History[i] = h
	}
}

func (h *History) CreateRows() {
	for _, r := range h.Rows {
		tmpr := r
		for {
			row := tmpr.CalculateNextRow()
			h.Rows = append(h.Rows, row)
			isZeroRow := true
			for _, v := range row.Row {
				if v != 0 {
					isZeroRow = false
					break
				}
			}
			if isZeroRow {
				break
			}
			tmpr = row
		}
	}
}

func (hr *HistoryRow) CalculateNextRow() HistoryRow {
	h := HistoryRow{Row: []int{}}
	for i := range hr.Row {
		if i == 0 {
			continue
		}
		diff := hr.Row[i] - hr.Row[i-1]
		h.Row = append(h.Row, diff)
		if i == len(hr.Row)-1 {
			break
		}
	}
	return h
}

func parseInput(input string) Oasis {
	o := Oasis{}
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		split := strings.Split(v, " ")
		h := History{}
		h.Rows = []HistoryRow{}
		row := []int{}
		for _, i := range split {
			k, _ := strconv.Atoi(i)
			row = append(row, k)
		}
		hrow := HistoryRow{Row: row}
		h.Rows = append(h.Rows, hrow)
		o.History = append(o.History, h)
	}
	return o
}
