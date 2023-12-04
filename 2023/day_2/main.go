package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	InputByRow         []string
	maxNumbersPerColor map[string]int
)

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	maxNumbersPerColor = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	InputByRow = strings.Split(string(file), "\n")
	colourAndValuePerRow := parseInput()
	approvedGames := solvePartOne(colourAndValuePerRow)
	sumPartOne := 0
	for _, val := range approvedGames {
		fmt.Printf("approved game %d \n", val)
		sumPartOne += val
	}
	fmt.Printf("Part 1 result: %d\n", sumPartOne)

	powers := solvePartTwo(colourAndValuePerRow)
	sumPowers := 0
	for _, val := range powers {
		sumPowers += val
	}

	fmt.Printf("Part 2 result: %d\n", sumPowers)
}

func parseInput() [][]map[string]int {
	rowsColourValues := [][]map[string]int{}
	for _, row := range InputByRow {
		if len(row) > 0 {
			parsed := parseRow(row)
			rowsColourValues = append(rowsColourValues, parsed)
		}
	}

	return rowsColourValues
}

func parseRow(row string) []map[string]int {
	valuesAndColoursOnRow := []map[string]int{}
	temp := strings.Split(row, ":")
	roundNumber := strings.Split(temp[0], " ")[1]
	strs := strings.Split(temp[1], ";")
	for _, str := range strs {
		valueAndColourFromSet := getValueAndColoursFromSet(str)
		valuesAndColoursOnRow = append(valuesAndColoursOnRow, valueAndColourFromSet)
		// for colour, value := range valueAndColourFromSet {
		// 	val, ok := valuesAndColoursOnRow[colour]
		// 	if ok {
		// 		valuesAndColoursOnRow[colour] = val + value
		// 	} else {
		// 		valuesAndColoursOnRow[colour] = value
		// 	}
		// }
	}
	fmt.Printf("Game %s, values: %v\n", roundNumber, valuesAndColoursOnRow)
	return valuesAndColoursOnRow
}

func getValueAndColoursFromSet(str string) map[string]int {
	valueAndColours := strings.Split(str, ", ")
	coloursWithValue := map[string]int{}
	for _, vac := range valueAndColours {
		vac = strings.Trim(vac, " ")
		splitted := strings.Split(vac, " ")
		value, _ := strconv.Atoi(strings.Trim(splitted[0], " "))
		colour := strings.Trim(splitted[1], " ")
		val, ok := coloursWithValue[colour]
		if ok {
			coloursWithValue[colour] = val + value
		} else {
			coloursWithValue[colour] = value
		}
	}
	return coloursWithValue
}

func solvePartOne(cavpr [][]map[string]int) []int {
	approvedGames := []int{}
	for idx, m := range cavpr {
		approved := verifyPartOne(m)
		if approved {
			approvedGames = append(approvedGames, idx+1)
		}
	}
	return approvedGames
}

func verifyPartOne(m []map[string]int) bool {
	for _, set := range m {
		for colour, value := range set {
			valMax, existsMax := maxNumbersPerColor[colour]
			if existsMax {
				if value > valMax {
					return false
				}
			}
		}
	}
	return true
}

func solvePartTwo(cavpr [][]map[string]int) []int {
	powers := []int{}
	for _, m := range cavpr {
		powers = append(powers, verifyPartTwo(m))
	}
	return powers
}

func verifyPartTwo(m []map[string]int) int {
	maxByColour := map[string]int{}
	power := 1
	for _, set := range m {
		for colour, value := range set {
			val, ok := maxByColour[colour]
			if ok {
				if value > val {
					maxByColour[colour] = value
				}
			} else {
				maxByColour[colour] = value
			}
		}
	}
	for _, value := range maxByColour {
		power = power * value
	}

	return power
}
