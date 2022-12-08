package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NodeValue struct {
	Name string
	Size int
	Type string
}

type Node struct {
	Data     NodeValue
	Children []*Node
	Parent   *Node
}

type LinkedList struct {
	Head *Node
}

var (
	CurrentNode *Node
	InputByRow  []string
	List        *LinkedList
)

func main() {
	file, err := os.ReadFile("./banana.txt")
	if err != nil {
		panic(err.Error())
	}
	List = &LinkedList{}
	nodeType := "directory"
	initNode := &Node{
		Data: NodeValue{
			Name: "/",
			Size: 0,
			Type: nodeType,
		},
	}
	CurrentNode = initNode
	List.Head = CurrentNode
	InputByRow = strings.Split(string(file), "\n")

	parseInput()
	solvePartOne()
	solvePartTwo()
}

func parseInput() {
	for k, str := range InputByRow {
		bytes := []byte(str)

		if len(bytes) > 0 {
			if bytes[0] == '$' {
				command, directory := parseCommand(str)
				lastIndex := executeCommand(command, *directory, k)
				if lastIndex > 0 {
					handleDirectoryContent(k+1, lastIndex)
				}
			}
		}

	}
}

func solvePartOne() {
	nodes := List.FindSizesSmallerThan(100000)
	totalSize := 0
	for _, v := range nodes {
		totalSize += v.Data.Size
	}
	fmt.Printf("Part One answer: %d\n", totalSize)
}

func solvePartTwo() {
	fullSize := 70000000
	updateSize := 30000000
	freeSpace := fullSize - List.Head.Data.Size
	neededSpace := updateSize - freeSpace
	nodes := List.FindSizesLargerThan(neededSpace)
	sort.SliceStable(nodes, func(a int, b int) bool {
		return nodes[a].Data.Size < nodes[b].Data.Size
	})
	fmt.Printf("THE ANSWER IS: %d\n", nodes[0].Data.Size)
}

func parseCommand(str string) (string, *string) {
	split := strings.Split(str, " ")
	var directory string

	if len(split) > 2 {
		directory = split[2]
	}

	return split[1], &directory
}

func executeCommand(command string, directory string, index int) int {
	if command == "cd" {
		changeDirectory(directory)
		return 0
	} else if command == "ls" {
		lastIndex := readListDirectory(index)
		return lastIndex
	}

	return 0
}

func changeDirectory(directory string) {
	node := CurrentNode.Parent
	if directory != ".." {
		node = List.Find(directory, CurrentNode)
	}
	CurrentNode = node
}

func (L *LinkedList) Find(name string, startNode *Node) *Node {
	firstNode := startNode
	if firstNode.Data.Type == "directory" && firstNode.Data.Name == name {
		return firstNode
	}
	found, searchNode := firstNode.traverseNodeTreeSearch(name)
	if !found {
		panic("Node could not be found")
	}

	return searchNode
}

func (N *Node) traverseNodeTreeSearch(name string) (bool, *Node) {
	if N.Data.Type == "directory" && N.Data.Name == name {
		return true, N
	}
	if len(N.Children) > 0 {
		for _, v := range N.Children {
			if v.Data.Type == "directory" && v.Data.Name == name {
				return true, v
			}
		}
	}
	return false, nil
}

func handleDirectoryContent(start int, end int) {
	for _, str := range InputByRow[start:end] {

		if str == "" {
			break
		}

		split := strings.Split(str, " ")

		if split[0] == "dir" {
			CurrentNode.AddNew(split[1], "directory", 0)
		} else {
			addSizeToCurrentNode(split)
		}
	}
}

func addSizeToCurrentNode(split []string) {
	size, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	name := split[1]

	n := CurrentNode.AddNew(name, "file", size)
	CurrentNode.Data.Size += n.Data.Size

	node := CurrentNode.Parent
	for node != nil {
		node.Data.Size += size
		node = node.Parent
	}
}

func readListDirectory(index int) int {
	newIndex := index
	if len(InputByRow) > index+1 {
		for _, str := range InputByRow[index+1:] {
			newIndex += 1
			byteText := []byte(str)
			if len(byteText) > 0 {
				if byteText[0] == '$' {
					break
				}
			}
		}
	}
	return newIndex
}

func (N *Node) AddNew(name string, nodeType string, size int) *Node {
	newNode := &Node{
		Data: NodeValue{
			Name: name,
			Size: size,
			Type: nodeType,
		},
	}
	newNode.Parent = N

	N.Children = append(N.Children, newNode)

	return newNode
}

func (L *LinkedList) FindSizesSmallerThan(size int) []*Node {
	firstNode := L.Head
	nodes := firstNode.traverseSizes(size, "smaller")
	return nodes
}

func (L *LinkedList) FindSizesLargerThan(size int) []*Node {
	firstNode := L.Head
	nodes := firstNode.traverseSizes(size, "larger")
	return nodes
}

func (L *LinkedList) PrintTree() {
	firstNode := L.Head
	RecursivePrint(firstNode)
}

func RecursivePrint(node *Node) {
	parentName := ""
	parentSize := 99999999
	if node.Parent != nil {
		parentName = node.Parent.Data.Name
		parentSize = node.Parent.Data.Size
	}
	if node.Data.Type == "directory" && node.Data.Size == 0 {
		fmt.Printf("Node %s with size %d in parent %s with size %d\n", node.Data.Name, node.Data.Size, parentName, parentSize)
	}
	if len(node.Children) > 0 {
		for _, v := range node.Children {
			RecursivePrint(v)
		}
	}
}

func (N *Node) traverseSizes(size int, largerorsmaller string) []*Node {
	tempNodes := []*Node{}
	if largerorsmaller == "larger" {
		if N.Data.Type == "directory" && N.Data.Size >= size {
			tempNodes = append(tempNodes, N)
		}
	} else {
		if N.Data.Type == "directory" && N.Data.Size <= size {
			tempNodes = append(tempNodes, N)
		}
	}

	if len(N.Children) > 0 {
		for _, n := range N.Children {
			if n.Data.Type == "directory" {
				newNodes := n.traverseSizes(size, largerorsmaller)
				tempNodes = append(tempNodes, newNodes...)
			}
		}
	}
	return tempNodes
}
