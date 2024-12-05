package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Printer struct {
	Rules      []Rule
	Prints     []Print
	SumAllowed int
}

type Rule struct {
	FirstNbr int
	LastNbr  int
}

type Print struct {
	PrintOrder []int
	Allowed    bool
	Solved     bool
}

func (p *Printer) CheckValidity(v int, pv []int) (int, bool) {
	for i, r := range p.Rules {
		if r.FirstNbr == v {
			if slices.Contains(pv, r.LastNbr) {
				return i, false
			}
		}
	}
	return -1, true
}

func (pr *Print) SolveNotAllowed(p *Printer) bool {
	localCopy := []int{}
	pv := []int{}
	failedRules := []Rule{}
	for _, po := range pr.PrintOrder {
		localCopy = append(localCopy, po)
		missIndex, ok := p.CheckValidity(po, pv)
		if !ok {
			failedRules = append(failedRules, p.Rules[missIndex])
		}
		pv = append(pv, po)
	}
	for _, r := range failedRules {
		firstIdx := -1
		lastVal := -1
		llc := []int{}
		for i, li := range localCopy {
			if li == r.FirstNbr {
				firstIdx = i
				llc = append(llc, li)
			} else if li == r.LastNbr {
				lastVal = li
			} else {
				if firstIdx != -1 {
					llc = append(llc, lastVal)
					firstIdx = -1
				}
				llc = append(llc, li)
			}
		}
		if firstIdx != -1 {
			llc = append(llc, lastVal)
			firstIdx = -1
		}
		lpr := Print{}
		lpr.PrintOrder = llc
		pr.PrintOrder = lpr.PrintOrder
		if lpr.CheckRow(p) {
			return true
		}
	}
	return false
}

func (p *Printer) HandleNotAllowed() {
	stillNotSolved := []Print{}
	for i, pr := range p.Prints {
		if !pr.Allowed {
			if !pr.SolveNotAllowed(p) {
				stillNotSolved = append(stillNotSolved, pr)
				pr.Solved = false
			} else {
				pr.Solved = true
			}
			p.Prints[i] = pr
		}
	}
	left := len(stillNotSolved)
	for {
		if left == 0 {
			break
		}
		for i, pr := range p.Prints {
			if !pr.Allowed && !pr.Solved {
				if pr.SolveNotAllowed(p) {
					pr.Solved = true
					left -= 1
				}
				p.Prints[i] = pr
			}
		}
	}

}

func (pr *Print) CheckRow(p *Printer) bool {
	pv := []int{}
	for _, po := range pr.PrintOrder {
		if _, ok := p.CheckValidity(po, pv); ok {
			pv = append(pv, po)
		} else {
			return false
		}
	}

	return true
}

func (pr *Print) FindMiddle() int {
	return pr.PrintOrder[int(math.Ceil(float64(len(pr.PrintOrder)/2)))]
}

func (p *Printer) FindRows() {
	for i, pr := range p.Prints {
		if pr.CheckRow(p) {
			pr.Allowed = true
			p.Prints[i] = pr
		}
	}
}

func (p *Printer) ParseInput(input string) {
	largeSplit := strings.Split(input, "\n\n")
	ruleSet := strings.Split(strings.TrimRight(largeSplit[0], "\n"), "\n")
	for _, split := range ruleSet {
		nums := strings.Split(split, "|")
		r := Rule{}
		val1, _ := strconv.Atoi(nums[0])
		val2, _ := strconv.Atoi(nums[1])
		r.FirstNbr = val1
		r.LastNbr = val2
		p.Rules = append(p.Rules, r)
	}
	printOrder := strings.Split(strings.TrimRight(largeSplit[1], "\n"), "\n")
	prints := []Print{}
	for _, split := range printOrder {
		pr := Print{}
		nums := strings.Split(split, ",")
		for _, n := range nums {
			nbr, _ := strconv.Atoi(n)
			pr.PrintOrder = append(pr.PrintOrder, nbr)
		}
		prints = append(prints, pr)
	}
	p.Prints = prints
}

func (p *Printer) Print() {
	fmt.Printf("Rules: \n")
	for _, r := range p.Rules {
		fmt.Printf("%d|%d\n", r.FirstNbr, r.LastNbr)
	}
	fmt.Printf("\n")
	fmt.Printf("Prints: \n")
	for _, pr := range p.Prints {
		fmt.Printf("%v\n", pr.PrintOrder)
	}
}

func (p *Printer) Solve(checkAllowed bool) {
	sum := 0
	for _, pr := range p.Prints {
		if pr.Allowed == checkAllowed {
			sum += pr.FindMiddle()
		}
	}
	p.SumAllowed = sum
}

func solvePartOne(input string) {
	p := Printer{}
	p.ParseInput(input)
	p.FindRows()
	// p.Print()
	p.Solve(true)
	fmt.Printf("Part 1: %d\n", p.SumAllowed)
}

func solvePartTwo(input string) {
	p := Printer{}
	p.ParseInput(input)
	p.FindRows()
	p.HandleNotAllowed()
	p.Solve(false)
	fmt.Printf("Part 2: %d\n", p.SumAllowed)
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}
