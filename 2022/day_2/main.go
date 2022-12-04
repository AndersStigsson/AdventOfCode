package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A X Rock 1
// B Y Paper 2
// C Z Scissors 3
// Win 6
// Draw 3
// Loss 0
func main() {
	file, _ := os.Open("./banana")

	defer file.Close()

	matrix := map[string]map[string]int{
		"A": {
			"X": 4,
			"Y": 8,
			"Z": 3,
		},
		"B": {
			"X": 1,
			"Y": 5,
			"Z": 9,
		},
		"C": {
			"X": 7,
			"Y": 2,
			"Z": 6,
		},
	}
	scanner := bufio.NewScanner(file)

	sumPoints := 0

	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")

		choice := getPartTwoChoice(splitText[0], splitText[1])
		sumPoints += matrix[splitText[0]][choice]
	}

	fmt.Println(sumPoints)
}

func getPartTwoChoice(a string, b string) string {
	matrixChoose := map[string]map[string]string{
		"A": {
			"X": "Z",
			"Y": "X",
			"Z": "Y",
		},
		"B": {
			"X": "X",
			"Y": "Y",
			"Z": "Z",
		},
		"C": {
			"X": "Y",
			"Y": "Z",
			"Z": "X",
		},
	}
	return matrixChoose[a][b]
}
