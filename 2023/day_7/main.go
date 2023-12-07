package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card struct {
	Key int
}

type Hand struct {
	Cards          []Card
	CardString     string
	Rank           int
	TotalValue     int
	Bid            int
	PokerHandValue int
}

type Game struct {
	Hands []Hand
}

var cardToInt = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func main() {
	fmt.Printf("Part One: %d\n", solvePartOne())
	fmt.Printf("Part Two: %d\n", solvePartTwo())
}

func solvePartOne() int {
	g := parseInput(input)
	g.CalculateValueOnHands(false)
	g.CalculateRanks()
	return g.CalculateResults()
}

func solvePartTwo() int {
	cardToInt['J'] = 1
	g := parseInput(input)
	g.CalculateValueOnHands(true)
	g.CalculateRanks()
	return g.CalculateResults()
}

func (g *Game) CalculateResults() int {
	sum := 0
	for _, v := range g.Hands {
		sum += v.Rank * v.Bid
	}

	return sum
}

func (g *Game) CalculateRanks() {
	sort.SliceStable(g.Hands, func(i, j int) bool {
		a := g.Hands[i]
		b := g.Hands[j]
		return a.PokerHandValue < b.PokerHandValue
	})

	sort.SliceStable(g.Hands, func(i, j int) bool {
		a := g.Hands[i]
		b := g.Hands[j]
		if a.PokerHandValue == b.PokerHandValue {
			return a.IsLower(b)
		}
		return false
	})

	for i, v := range g.Hands {
		v.Rank = i + 1
		g.Hands[i] = v
	}
}

func (g *Game) CalculateValueOnHands(roundTwo bool) {
	for i, h := range g.Hands {
		h.CalculateValueOnHand(roundTwo)
		g.Hands[i] = h
	}
}

func (h *Hand) CalculateValueOnHand(includeJokers bool) {
	aoc := map[int]int{}
	numJokers := 0
	for _, c := range h.Cards {
		if c.Key == cardToInt['J'] {
			numJokers++
			continue
		}
		aoc[c.Key]++
	}

	largestAmount := 0
	secondLargestAmount := 0

	for i, v := range aoc {
		if includeJokers && i == cardToInt['J'] {
			continue
		}
		if v >= largestAmount {
			if secondLargestAmount < largestAmount {
				secondLargestAmount = largestAmount
			}
			largestAmount = v
		} else if v >= secondLargestAmount {
			secondLargestAmount = v
		}
	}

	if includeJokers && numJokers > 0 {
		largestAmount, secondLargestAmount = getHighestAndLowestWithJokers(largestAmount, secondLargestAmount, numJokers)
	}
	pokerHandValue := getPokerHandValue(largestAmount, secondLargestAmount)
	h.PokerHandValue = pokerHandValue
}

func (h *Hand) IsLower(h2 Hand) bool {
	for i, v := range h.Cards {
		if v.Key < h2.Cards[i].Key {
			return true
		} else if v.Key > h2.Cards[i].Key {
			return false
		}
	}

	return false
}

func getHighestAndLowestWithJokers(la int, sa int, jokers int) (int, int) {
	largestAmount := la + jokers

	if largestAmount > 3 {
		return largestAmount, 0
	}

	return largestAmount, sa
}

func getPokerHandValue(la int, sa int) int {
	pokerHandValue := la
	if pokerHandValue == 1 {
		return 1
	}
	if sa == 2 {
		return pokerHandValue*2 + 1
	}

	return pokerHandValue * 2
}

func parseInput(input string) Game {
	g := Game{}
	input = strings.TrimRight(input, "\n")
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		h := Hand{}
		split := strings.Split(v, " ")
		h.CardString = split[0]
		for _, c := range split[0] {
			card := Card{}
			if val, ok := cardToInt[c]; ok {
				card.Key = val
			}
			h.Cards = append(h.Cards, card)
		}
		val, _ := strconv.Atoi(split[1])
		h.Bid = val
		g.Hands = append(g.Hands, h)
	}
	return g
}
