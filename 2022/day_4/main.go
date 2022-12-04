package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./banana")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, ",")
		firstPart := strings.Split(splitText[0], "-")
		secondPart := strings.Split(splitText[1], "-")
		lowerThresholdFirstPart, _ := strconv.Atoi(firstPart[0])
		upperThresholdFirstPart, _ := strconv.Atoi(firstPart[1])

		lowerThresholdSecondPart, _ := strconv.Atoi(secondPart[0])
		upperThresholdSecondPart, _ := strconv.Atoi(secondPart[1])

		if lowerThresholdSecondPart >= lowerThresholdFirstPart &&
			upperThresholdSecondPart <= upperThresholdFirstPart {
			total += 1
		} else if lowerThresholdFirstPart >= lowerThresholdSecondPart &&
			upperThresholdFirstPart <= upperThresholdSecondPart {
			total += 1
		}
	}
	fmt.Printf("Part 1: %d\n", total)
	partTwo()
}

func partTwo() {
	file, _ := os.Open("./banana")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, ",")
		firstPart := strings.Split(splitText[0], "-")
		secondPart := strings.Split(splitText[1], "-")

		lowerThresholdFirstPart, _ := strconv.Atoi(firstPart[0])
		upperThresholdFirstPart, _ := strconv.Atoi(firstPart[1])

		lowerThresholdSecondPart, _ := strconv.Atoi(secondPart[0])
		upperThresholdSecondPart, _ := strconv.Atoi(secondPart[1])

		// if second lower threshold is between lower and upper on first part
		if lowerThresholdSecondPart >= lowerThresholdFirstPart &&
			lowerThresholdSecondPart <= upperThresholdFirstPart {
			total += 1
			// if first lower threshold is between lower and upper on second part
		} else if lowerThresholdFirstPart >= lowerThresholdSecondPart &&
			lowerThresholdFirstPart <= upperThresholdSecondPart {
			total += 1
		} else if upperThresholdSecondPart >= lowerThresholdFirstPart &&
			upperThresholdSecondPart <= upperThresholdFirstPart {
			total += 1
		} else if upperThresholdFirstPart >= lowerThresholdSecondPart &&
			upperThresholdFirstPart <= upperThresholdSecondPart {
			total += 1
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}
