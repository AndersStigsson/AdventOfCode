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

type CardList struct {
	Cards []Card
}

type Card struct {
	Numbers        []int
	WinningNumbers []int
	WinPoints      int
	ExtraCards     []Card
	Index          int
}

func main() {
	solvePartOne(input)
	solvePartTwo(input)
}

func solvePartOne(input string) {
	cl := CardList{}
	cl.ParseInput(input)
	cl.CalculateWinPoints()
	fmt.Printf("Answer part 1: %d\n", cl.SolvePartOne())
}

func solvePartTwo(input string) {
	cl := CardList{}
	cl.ParseInput(input)
	fmt.Printf("Answer part 2: %d\n", cl.SolvePartTwoRecursion())
}

func (cl *CardList) SolvePartOne() int {
	sum := 0
	for _, c := range cl.Cards {
		sum += c.WinPoints
	}
	return sum
}

func (cl *CardList) SolvePartTwoRecursion() int {
	for idx, c := range cl.Cards {
		c.CalculateExtraCardsRecursive(*cl)
		cl.Cards[idx] = c
	}

	return cl.RecurseOverExtraCards()
}

func (cl *CardList) RecurseOverExtraCards() int {
	sum := 0
	for _, c := range cl.Cards {
		sum = c.RecurseOverCards(sum)
	}

	sum += len(cl.Cards)

	return sum
}

func (cl *CardList) ParseInput(input string) {
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for idx, str := range splitted {
		splitstr := strings.Split(str, ":")
		c := parseCard(splitstr[1])
		c.Index = idx
		cl.Cards = append(cl.Cards, c)
	}
}

func (cl *CardList) CalculateWinPoints() {
	for idx, c := range cl.Cards {
		winningPoints := c.CalculateWinPoints()
		c.WinPoints = 0
		if len(winningPoints) > 0 {
			c.WinPoints = int(math.Pow(float64(2), float64(len(winningPoints)-1)))
		}
		cl.Cards[idx] = c
	}
}

func (c *Card) RecurseOverCards(length int) int {
	for _, ec := range c.ExtraCards {
		length = ec.RecurseOverCards(length) + 1
	}

	return length
}

func (c *Card) CalculateWinPoints() []int {
	includedPoints := []int{}
	for _, v := range c.Numbers {
		if contains(c.WinningNumbers, v) {
			includedPoints = append(includedPoints, v)
		}
	}

	return includedPoints
}

func (c *Card) CalculateExtraCardsRecursive(cl CardList) {
	cardNbr := c.Index + 1
	for _, v := range c.Numbers {
		if contains(c.WinningNumbers, v) {
			extraCard := Card{}
			extraCard.WinningNumbers = cl.Cards[cardNbr].WinningNumbers
			extraCard.Numbers = cl.Cards[cardNbr].Numbers
			extraCard.Index = cl.Cards[cardNbr].Index
			extraCard.CalculateExtraCardsRecursive(cl)
			c.ExtraCards = append(c.ExtraCards, extraCard)
			cardNbr += 1
		}
	}
}

func contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func parseCard(str string) Card {
	c := Card{}
	split := strings.Split(str, "|")
	for idx, s := range split {
		splitted := strings.Split(s, " ")
		for _, t := range splitted {
			if t == "" {
				continue
			}
			intval, _ := strconv.Atoi(t)
			if idx == 0 {
				c.WinningNumbers = append(c.WinningNumbers, intval)
			} else {
				c.Numbers = append(c.Numbers, intval)
			}
		}
	}
	return c
}
