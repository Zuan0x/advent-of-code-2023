package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	task1(string(file))
	//task2(rows)
}

type tNode struct {
	left  string
	right string
}

func task1(file string) {
	rows := strings.Split(file, "\n\n")

	instructions := rows[0]
	nodeList := strings.Split(rows[1], "\n")

	fmt.Println(instructions)
	fmt.Println(nodeList)

	nodes := make(map[string]tNode)

	for _, node := range nodeList {

		nodeParts := strings.Split(node, " = ")
		nodeKey := nodeParts[0]

		trimmedString := strings.Trim(nodeParts[1], "()")

		// Split the string by comma and trim spaces
		destinations := strings.Split(trimmedString, ",")

		fmt.Println(node, nodeKey, destinations)
		nodes[nodeKey] = tNode{
			left:  strings.Trim(destinations[0], " "),
			right: strings.Trim(destinations[1], " "),
		}
	}

	pos := "AAA"
	idx := 0

	for {
		direction := string(instructions[idx%len(instructions)])

		if direction == "R" {
			pos = nodes[pos].right
		} else if direction == "L" {
			pos = nodes[pos].left
		} else {
			fmt.Println("Error")
		}

		idx++

		if pos == "ZZZ" {
			break
		}
	}

	fmt.Println(idx)

}
