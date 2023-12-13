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
	count := 0
	startCount := 0
	for i := 0; i < len(rows); i++ {
		fmt.Println(i)

		splitString := strings.Split(rows[i], " ")
		rowInt := make([]int, 0)

		//Check if first and last values are the same
		if splitString[0] != splitString[len(splitString)-1] {
			//Convert string array to int array
			rowInt = convertStringArrayToIntArray(splitString)
			//Print difference in values
			finalDif, startDif := findDifference(rowInt, 0, 0)
			fmt.Println("Final dif:", finalDif, startDif)
			count += finalDif
			startCount += startDif
		}
	}
	fmt.Println(count, startCount)
}

func task2(file []string) {
}

func findDifference(values []int, dif int, startDif int) (int, int) {
	allSame := allSameInt(values)
	if allSame == false {
		differences := make([]int, 0)
		for i := 1; i < len(values); i++ {
			differences = append(differences, values[i]-values[i-1])
		}
		fmt.Println(values)
		dif, startDif = findDifference(differences, dif, startDif)
		fmt.Println("dif:", dif, startDif)

	} else {
		fmt.Println(values)
		return values[0], values[0]
	}

	return values[len(values)-1] + dif, values[0] - startDif
}

func allSameInt(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func convertStringArrayToIntArray(stringArray []string) []int {
	var intArray []int

	for _, str := range stringArray {
		// Convert each string element to an int
		num, err := strconv.Atoi(str)
		if err != nil {
			// Handle the error (e.g., invalid conversion)
			return nil
		}

		// Append the converted int to the int array
		intArray = append(intArray, num)
	}

	return intArray
}
