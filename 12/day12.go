package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Split the file into lines
	rows := strings.Split(string(file), "\n")

	task1(rows)
	//task2(string(file))
}

func task1(rows []string) {
	result := 0
	for i := 0; i < len(rows); i++ {
		//Split the row into order (???.###) and groups (1,1,3)
		splitRow := strings.Split(rows[i], " ")
		order := splitRow[0]
		groups := strings.Split(splitRow[1], ",")
		groupsInt, err := convertStringArrayToIntArray(groups)
		expected := groupsInt[0]
		fmt.Println(order, expected, err)
		count := back(order, groupsInt, groupsInt)
		result += count
	}
	fmt.Println(result)
}

func back(order string, expected []int, groups []int) int {

	index := strings.Index(order, "?")

	//If there are no ? in the order, check if the order is valid
	if index == -1 {
		if valid(order, expected) {
			return 1
		} else {
			return 0
		}
	}
	//If there is a ? in the order, replace it with # and .
	order = order[:index] + "#" + order[index+1:]
	a := back(order, expected, groups)
	order = order[:index] + "." + order[index+1:]
	b := back(order, expected, groups)
	order = order[:index] + "?" + order[index+1:]

	//Return the result of the two recursive calls
	return a + b
}

func valid(order string, expected []int) bool {
	current := 0
	actual := make([]int, 0)

	for c := 0; c < len(order); c++ {
		if (order[c] == '.' || order[c] == '?') && (current > 0) {
			actual = append(actual, current)
			current = 0
		}
		if order[c] == '#' {
			current++
		}
	}

	if current > 0 {

		actual = append(actual, current)
	}
	return IntArrayEquals(actual, expected)
}

func convertStringArrayToIntArray(stringArray []string) ([]int, error) {
	var intArray []int

	for _, str := range stringArray {
		// Convert each string element to an int
		num, err := strconv.Atoi(str)
		if err != nil {
			// Handle the error (e.g., invalid conversion)
			return nil, err
		}

		// Append the converted int to the int array
		intArray = append(intArray, num)
	}

	return intArray, nil
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func task2(file []string) {
}
