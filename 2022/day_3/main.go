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
	result := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		part1 := text[0 : len(text)/2]
		part2 := text[len(text)/2:]
		for _, c := range part2 {
			if strings.Contains(part1, string(c)) {
				total += strings.Index(result, string(c)) + 1
				break
			}
		}
	}
	fmt.Printf("Part 1: %d \n", total)
	partTwo()
}

func partTwo() {
	file, _ := os.Open("./banana")
	result := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	scanner := bufio.NewScanner(file)

	idx := 0
	total := 0
	texts := []string{}
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
		idx += 1
		if idx%3 == 0 {
			inOneAndTwo := make(map[string]bool)
			for _, c := range texts[1] {
				if strings.Contains(texts[0], string(c)) {
					inOneAndTwo[string(c)] = true
				}
			}
			for _, c := range texts[2] {
				if inOneAndTwo[string(c)] {
					total += strings.Index(result, string(c)) + 1
					idx = 0
					texts = []string{}
					break
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}
