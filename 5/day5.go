package main

import (
	"fmt"
	"math"
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
	maps := strings.Split(string(file), "\n\n")

	task1(maps)
}

func task1(maps []string) {
	seeds := strings.Split(strings.Split(maps[0], ": ")[1], " ")
	//Convert seeds to int
	seedsInt, err := convertStringArrayToIntArray(seeds)
	if err != nil {
		// ... handle error
		panic(err)
	}

	fmt.Println(seedsInt)

	maps = maps[1:]

	for i := 0; i < len(maps); i++ {
		fmt.Println("Block ", i)

		ranges := make([][]string, 0)
		lines := strings.Split(maps[i], "\n")[1:]
		for j := 0; j < len(lines); j++ {
			fmt.Println(lines[j])
			ranges = append(ranges, strings.Split(lines[j], " "))
		}

		new := make([]int, 0)

		for j := 0; j < len(seedsInt); j++ {
			seed := seedsInt[j]
			valAdded := false
			for _, rangeValues := range ranges {
				dest := convertStrToInt(rangeValues[0])
				source := convertStrToInt(rangeValues[1])
				gap := convertStrToInt(rangeValues[2])
				if source <= seed && seed < source+gap {
					fmt.Println(seed, source, gap)
					new = append(new, seed-source+dest)
					valAdded = true
					break
				}
				//new = append(new, seed-source+gap)
			}
			if !valAdded {
				new = append(new, seed)
			}
		}
		seedsInt = new
		fmt.Println(seedsInt)
	}

	fmt.Println(findMinValue(seedsInt))
}

func convertStrToInt(str string) int {
	output, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return output
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

func findMinValue(numbers []int) int {
	// Initialize minValue with the maximum possible integer value
	minValue := math.MaxInt

	// Iterate through the elements and find the minimum
	for _, num := range numbers {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}
