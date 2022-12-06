package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("./banan.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []int
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			result = append(result, total)
			total = 0
		}
		intval, _ := strconv.Atoi(text)
		total += intval
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[j] < result[i]
	})
	fmt.Printf("Solution to first question: %d\n", result[0])

	sum := 0
	for i := 0; i < 3; i++ {
		sum += result[i]
	}
	fmt.Printf("Solution to second question: %d\n", sum)
}
