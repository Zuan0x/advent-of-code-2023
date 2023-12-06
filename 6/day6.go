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

	//task1(rows)
	task2(rows)
}

func task1(rows []string) {
	times, err := convertStringArrayToIntArray(strings.Split(strings.Split(rows[0], ":")[1], " "))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	distances, err := convertStringArrayToIntArray(strings.Split(strings.Split(rows[1], ":")[1], " "))

	total := 1

	//For each race
	for i := 0; i < len(times); i++ {
		wins := 0
		//Loop through button hold times
		for j := 0; j < times[i]; j++ {
			//Loop through each distance

			fmt.Println(j, j*(times[i]-j))
			distanceTravelled := j * (times[i] - j)

			if distanceTravelled > distances[i] {
				wins++
			}
		}
		total *= wins
	}

	fmt.Println(total)
}

func task2(rows []string) {
	times := strings.Split(strings.Split(rows[0], ":")[1], " ")
	//Trim spaces
	for i := 0; i < len(times); i++ {
		times[i] = strings.Trim(times[i], " ")
	}
	joinedTime := strings.Join(times, "")
	joinedTimeInt := convertStrToInt(joinedTime)
	distances := strings.Split(strings.Split(rows[1], ":")[1], " ")
	//Trim spaces
	for i := 0; i < len(distances); i++ {
		distances[i] = strings.Trim(distances[i], " ")
	}
	joinedDistance := strings.Join(distances, "")
	joinedDistanceInt := convertStrToInt(joinedDistance)

	total := 1

	start := 0
	//Loop through button hold times until you find a winning time
	for j := 0; j < joinedTimeInt; j++ {
		//Loop through each distance

		fmt.Println(j, j*(joinedTimeInt-j))
		distanceTravelled := j * (joinedTimeInt - j)

		if distanceTravelled > joinedDistanceInt {
			fmt.Println("WIN")
			start = j - 1
			break
		}
	}

	end := 0
	//Loop through button hold times until you find a winning time
	for j := joinedTimeInt; j > 0; j-- {
		//Loop through each distance

		fmt.Println(j, j*(joinedTimeInt-j))
		distanceTravelled := j * (joinedTimeInt - j)

		if distanceTravelled > joinedDistanceInt {
			fmt.Println("WIN END")
			end = j
			break
		}
	}

	total = end - start

	fmt.Println(start, end, total)
}

func convertStrToInt(str string) int {
	output, err := strconv.Atoi(strings.Trim(str, " "))
	if err != nil {
		// ... handle error
		panic(err)
	}
	return output
}

func convertStringArrayToIntArray(stringArray []string) ([]int, error) {
	var intArray []int

	for _, str := range stringArray {
		if str != "" {
			// Convert each string element to an int
			num, err := strconv.Atoi(strings.Trim(str, " "))
			if err != nil {
				// Handle the error (e.g., invalid conversion)
				return nil, err
			}

			// Append the converted int to the int array
			intArray = append(intArray, num)
		}
	}

	return intArray, nil
}
