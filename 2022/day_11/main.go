package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed banana.txt
var input string

type Value struct {
	Type  string
	Value string
}

type Operation struct {
	Type        string
	FirstValue  Value
	SecondValue Value
}

type Action struct {
	Recipient int
}

type Test struct {
	Value int
	True  Action
	False Action
}

type Item struct {
	Value int
	Id    string
}

type Monkey struct {
	Items          []Item
	Operation      Operation
	Test           Test
	InspectedItems int
}

type State struct {
	Monkeys         []Monkey
	Items           []Item
	MaxValue        int
	TogetherDivisor int
}

func parseStartingItems(str string, index int) []Item {
	returnValues := []Item{}
	re := regexp.MustCompile(`Starting items: (.*)`)
	matches := re.FindStringSubmatch(str)
	values := strings.Split(matches[1], ", ")
	for i, v := range values {
		intval, _ := strconv.Atoi(v)
		returnValues = append(returnValues, Item{Value: intval, Id: fmt.Sprintf("Monkey %d, item %d", index, i)})
	}

	return returnValues
}

func parseOperation(str string) Operation {
	operation := Operation{}
	re := regexp.MustCompile(`Operation: new (.*) (.) (.*)$`)
	matches := re.FindStringSubmatch(str)
	_, err := strconv.Atoi(matches[1])
	if err != nil {
		operation.FirstValue.Type = "string"
	} else {
		operation.FirstValue.Type = "int"
	}
	operation.FirstValue.Value = matches[1]
	_, err = strconv.Atoi(matches[3])
	if err != nil {
		operation.SecondValue.Type = "string"
	} else {
		operation.SecondValue.Type = "int"
	}
	operation.SecondValue.Value = matches[3]
	operation.Type = matches[2]
	return operation
}

func parseTest(split []string, index int) Test {
	test := Test{}
	patterns := []string{
		`Test: divisible by (\d+)$`,
		`If true: throw to monkey (\d+)$`,
		`If false: throw to monkey (\d+)$`,
	}
	filters := make([]*regexp.Regexp, len(patterns))

	for idx, pattern := range patterns {
		filters[idx] = regexp.MustCompile(pattern)
	}
	output := make(map[int]int, 3)
	for k, v := range split[index:] {
		if len(v) == 0 {
			continue
		}
		matches := filters[k].FindStringSubmatch(v)
		intval, _ := strconv.Atoi(matches[1])
		output[k] = intval
	}
	test.Value = output[0]
	test.True.Recipient = output[1]
	test.False.Recipient = output[2]

	return test
}

func parseMonkey(str string) Monkey {
	split := strings.Split(str, "\n")
	monkey := Monkey{InspectedItems: 0}
	for i, v := range split {
		if strings.Contains(v, "Starting") {
			monkey.Items = parseStartingItems(v, i)
		}
		if strings.Contains(v, "Operation") {
			monkey.Operation = parseOperation(v)
		}
		if strings.Contains(v, "Test") {
			monkey.Test = parseTest(split, i)
		}
	}
	return monkey
}

func (State *State) parseInput() {
	splits := strings.Split(input, "\n\n")
	for _, s := range splits {
		State.Monkeys = append(State.Monkeys, parseMonkey(s))
	}
}

func (State *State) DoOperationOnItem(item Item, divideWorryLevel bool, index int) int {
	operation := State.Monkeys[index].Operation
	firstValue := operation.FirstValue
	secondValue := operation.SecondValue
	var firstValueInt int
	if firstValue.Type == "int" {
		intval, _ := strconv.Atoi(firstValue.Value)
		firstValueInt = intval
	} else {
		firstValueInt = item.Value
	}
	var secondValueInt int
	if secondValue.Type == "int" {
		intval, _ := strconv.Atoi(secondValue.Value)
		secondValueInt = intval
	} else {
		secondValueInt = item.Value
	}
	worryLevel := item.Value
	switch operation.Type {
	case "*":
		worryLevel = firstValueInt * secondValueInt
	case "+":
		worryLevel = firstValueInt + secondValueInt
	case "-":
		worryLevel = firstValueInt - secondValueInt
	}
	State.Monkeys[index].InspectedItems++
	if divideWorryLevel {
		if worryLevel > State.MaxValue {
			State.MaxValue = worryLevel
		}
		return worryLevel / 3
	}
	if worryLevel > State.MaxValue {
		State.MaxValue = worryLevel
	}
	// if worryLevel >= math.MaxInt32 {
	// 	for _, v := range State.Items {
	// 		if v.Id == item.Id {
	// 			return v.Value
	// 		}
	// 	}
	// }
	return worryLevel % State.TogetherDivisor
}

func (State *State) MoveItemToNextMonkey(item Item, index int) {
	test := State.Monkeys[index].Test
	if item.Value%test.Value == 0 {
		State.Monkeys[test.True.Recipient].Items = append(State.Monkeys[test.True.Recipient].Items, item)
	} else {
		State.Monkeys[test.False.Recipient].Items = append(State.Monkeys[test.False.Recipient].Items, item)
	}
}

func (State *State) HandleAction(index int, divideWorryLevel bool) {
	for k := range State.Monkeys[index].Items {
		State.Monkeys[index].Items[k].Value = State.DoOperationOnItem(State.Monkeys[index].Items[k], divideWorryLevel, index)
		State.MoveItemToNextMonkey(State.Monkeys[index].Items[k], index)
	}
	State.Monkeys[index].Items = []Item{}
}

func (State *State) resetItems(index int) {
	for i, k := range State.Monkeys[index].Items {
		for _, v := range State.Items {
			if k.Id == v.Id {
				State.Monkeys[index].Items[i].Value = v.Value
			}
		}
	}
}

func (State *State) calculateDivisor() {
	totalDivisor := 1
	for _, v := range State.Monkeys {
		totalDivisor *= v.Test.Value
	}
	State.TogetherDivisor = totalDivisor
}

func (State *State) doStuff(divideWorryLevel bool, rounds int) {
	for i := 0; i < rounds; i++ {
		// if i > 10 && (i%50 == 0) && !divideWorryLevel {
		// 	for k := range State.Monkeys {
		// 		State.resetItems(k)
		// 	}
		// }
		for k := range State.Monkeys {
			State.HandleAction(k, divideWorryLevel)
		}
	}
}

func (State *State) FindMostInspectingMonkeys() []Monkey {
	sort.SliceStable(State.Monkeys, func(a int, b int) bool {
		return State.Monkeys[a].InspectedItems > State.Monkeys[b].InspectedItems
	})
	return State.Monkeys[:2]
}

func (State *State) PrintMonkeys() {
	for k, v := range State.Monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", k, v.InspectedItems)
	}
}

func partOne() {
	state := State{}
	state.parseInput()
	state.doStuff(true, 20)
	state.PrintMonkeys()
	mostInspectingMonkeys := state.FindMostInspectingMonkeys()
	totalValue := 1
	for _, v := range mostInspectingMonkeys {
		totalValue *= v.InspectedItems
	}
	fmt.Printf("Total InspectionValue: %d\n", totalValue)
	fmt.Printf("State max: %d\n", state.MaxValue)
}

func partTwo() {
	state := State{}
	state.parseInput()
	for _, v := range state.Monkeys {
		state.Items = append(state.Items, v.Items...)
	}
	state.calculateDivisor()
	state.doStuff(false, 10000)
	state.PrintMonkeys()
	mostInspectingMonkeys := state.FindMostInspectingMonkeys()
	totalValue := 1
	for _, v := range mostInspectingMonkeys {
		totalValue *= v.InspectedItems
	}
	fmt.Printf("Total InspectionValue: %d\n", totalValue)
	fmt.Printf("State max: %d", state.MaxValue)
}

func main() {
	partOne()
	partTwo()
}
