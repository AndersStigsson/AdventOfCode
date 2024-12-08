package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type EquationList struct {
	Equations []Equation
	TotalSum  int
}

type Equation struct {
	Total        int
	Numbers      []int
	Operator     string
	Solvable     bool
	CurrentTotal int
}

func (eq *Equation) TestSmallerPart() bool {
	if len(eq.Numbers) == 1 {
		if eq.CurrentTotal+eq.Numbers[0] == eq.Total {
			return true
		}
		if eq.CurrentTotal*eq.Numbers[0] == eq.Total {
			return true
		}
		lt := fmt.Sprintf("%d%d", eq.CurrentTotal, eq.Numbers[0])
		val, _ := strconv.Atoi(lt)
		if val == eq.Total {
			return true
		}
		return false
	}

	localTotal := eq.CurrentTotal + eq.Numbers[0]
	eqq := Equation{Total: eq.Total, CurrentTotal: localTotal, Numbers: eq.Numbers[1:]}
	if eqq.TestSmallerPart() {
		eq.Operator = eqq.Operator
		eq.Operator += "+"
		return true
	}

	localTotal = eq.CurrentTotal * eq.Numbers[0]
	eqq = Equation{Total: eq.Total, CurrentTotal: localTotal, Numbers: eq.Numbers[1:]}
	if eqq.TestSmallerPart() {
		eq.Operator = eqq.Operator
		eq.Operator += "*"
		return true
	}

	lt := fmt.Sprintf("%d%d", eq.CurrentTotal, eq.Numbers[0])
	val, _ := strconv.Atoi(lt)
	localTotal = val
	eqq = Equation{Total: eq.Total, CurrentTotal: localTotal, Numbers: eq.Numbers[1:]}
	if eqq.TestSmallerPart() {
		eq.Operator = eqq.Operator
		eq.Operator += "||"
		return true
	}

	return false
}

func (eq *Equation) TestEquation() bool {
	// firstNbr := eq.Numbers[0]
	// localTotal := firstNbr
	if eq.TestSmallerPart() {
		return true
	}
	// eqq := Equation{Total: eq.Total - localTotal, Numbers: eq.Numbers[1:]}
	// if eqq.TestSmallerPart() {
	// 	eq.Operator += "+"
	// 	eq.Operator = eqq.Operator
	// 	return true
	// }

	// eqq = Equation{Total: eq.Total / localTotal, Numbers: eq.Numbers[1:]}
	// if eqq.TestSmallerPart() {
	// 	eq.Operator += "*"
	// 	eq.Operator = eqq.Operator
	// 	return true
	// }

	return false
}

func (e *EquationList) FindSolution() {
	for i, eq := range e.Equations {
		if eq.TestEquation() {
			fmt.Printf("Eq %d, %d: %v solvable with operator %s\n", i, eq.Total, eq.Numbers, eq.Operator)
			eq.Solvable = true
			e.Equations[i] = eq
		}
	}
}

func (e *EquationList) Solve() {
	for _, eq := range e.Equations {
		if eq.Solvable {
			e.TotalSum += eq.Total
		}
	}
}

func (e *EquationList) ParseInput(input string) {
	splitted := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for _, split := range splitted {
		eq := Equation{}
		s := strings.Split(split, ":")
		tot, _ := strconv.Atoi(s[0])
		eq.Total = tot
		s = strings.Split(s[1][1:], " ")
		for _, ss := range s {
			val, _ := strconv.Atoi(ss)
			eq.Numbers = append(eq.Numbers, val)
		}
		e.Equations = append(e.Equations, eq)
	}
}

func (e *EquationList) Print() {
	for _, eq := range e.Equations {
		fmt.Printf("%d: %v\n", eq.Total, eq.Numbers)
	}
}

func solvePartOne(input string) {
	e := EquationList{}
	e.ParseInput(input)
	e.Print()
	e.FindSolution()
	e.Solve()
	fmt.Printf("Part 1: %d\n", e.TotalSum)
}

func main() {
	solvePartOne(input)
}
