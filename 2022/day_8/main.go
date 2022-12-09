package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	Height       int
	Visible      bool
	VisibleTrees int
}

type Forest struct {
	Trees [][]Tree
}

//go:embed banana.txt
var input string

func parseInput() Forest {
	forest := Forest{}
	for k, r := range strings.Split(input, "\n") {
		if len(r) == 0 {
			continue
		}
		if len(forest.Trees) <= k {
			forest.Trees = append(forest.Trees, []Tree{})
		}

		for _, c := range r {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err.Error())
			}
			forest.Trees[k] = append(forest.Trees[k], Tree{Height: height, Visible: false, VisibleTrees: 0})
		}
	}
	return forest
}

func (f *Forest) checkVisibilityOfTrees() {
	for i, v := range f.Trees {
		for k := range v {
			if f.checkIfFirstLine(i, k) {
				v[k].Visible = true
				continue
			}

			visible, totalVisible := f.checkNeighbouringTrees(i, k, v[k])
			v[k].VisibleTrees = totalVisible
			if visible {
				v[k].Visible = true
			}
		}
	}
}

func (f *Forest) checkIfFirstLine(i int, k int) bool {
	returnvalue := i == 0 ||
		i == len(f.Trees[i])-1 ||
		k == 0 ||
		k == len(f.Trees)-1
	return returnvalue
}

func (f *Forest) checkNeighbouringTrees(i int, k int, tree Tree) (bool, int) {
	// fmt.Printf("Checking if tree %d at position [%d, %d] is visible\n", tree.Height, i, k)

	up := true
	down := true
	left := true
	right := true
	totalVisible := 1

	index := 0
	for j := i + 1; j < len(f.Trees[i]); j++ {
		index += 1
		if f.Trees[j][k].Height >= tree.Height {
			down = false
			break
		}
	}
	totalVisible *= index
	index = 0
	for j := k + 1; j < len(f.Trees); j++ {
		index += 1
		if f.Trees[i][j].Height >= tree.Height {
			right = false
			break
		}
	}
	totalVisible *= index

	index = 0
	for j := i - 1; j >= 0; j-- {
		index += 1
		if f.Trees[j][k].Height >= tree.Height {
			up = false
			break
		}
	}

	totalVisible *= index

	index = 0
	for j := k - 1; j >= 0; j-- {
		index += 1
		if f.Trees[i][j].Height >= tree.Height {
			left = false
			break
		}
	}
	totalVisible *= index
	return (down || up || right || left), totalVisible
}

func (f *Forest) countVisibleTrees() int {
	totalVisible := 0
	for _, v := range f.Trees {
		for _, tree := range v {
			if tree.Visible {
				totalVisible += 1
			}
		}
	}
	return totalVisible
}

func (f *Forest) findHighestScenicScore() int {
	highestScore := 0
	for _, v := range f.Trees {
		for _, tree := range v {
			if tree.VisibleTrees > highestScore {
				highestScore = tree.VisibleTrees
			}
		}
	}
	return highestScore
}

func (f *Forest) printNumberOfVisibleTreesFromEachTree() {
	for _, v := range f.Trees {
		for _, tree := range v {
			fmt.Printf("%d ", tree.VisibleTrees)
		}
		fmt.Printf("\n")
	}
}

func (f *Forest) printVisibleTrees() {
	for _, v := range f.Trees {
		for _, tree := range v {
			if tree.Visible {
				fmt.Print("1")
			} else {
				fmt.Printf("0")
			}
		}
		fmt.Printf("\n")
	}
}

func partOne() {
	f := parseInput()
	f.checkVisibilityOfTrees()
	visibleTrees := f.countVisibleTrees()
	fmt.Printf("Total visible trees: %d\n", visibleTrees)
	// f.printVisibleTrees()
}

func partTwo() {
	f := parseInput()
	f.checkVisibilityOfTrees()
	highestScenicScore := f.findHighestScenicScore()
	fmt.Printf("Highest scenic score: %d\n", highestScenicScore)
	// f.printNumberOfVisibleTreesFromEachTree()
}

func main() {
	partOne()
	partTwo()
}
