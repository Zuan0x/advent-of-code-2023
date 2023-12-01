package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	task1()
}

func task1() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// Split the file into lines
	lines := strings.Split(string(file), "\n")

	count := 0

	// using for loop
	for i := 0; i < len(lines); i++ {
		//Loop through each character in the line
		//Make a new string only containing the first and last digit
		newString := ""
		for _, char := range lines[i] {
			// Check if the character is a digit
			if slices.Contains(digits, char) {
				//Append newString with only the digits
				newString += string(char)
			}
		}

		//If the length of the new string is 2, then print the line
		if len(newString) != 2 {
			//If the length of the new string is greater than 2, remove all but the first and last character
			newString = string(newString[0]) + string(newString[len(newString)-1])
		}

		fmt.Println(lines[i], newString)

		// string to int
		i, err := strconv.Atoi(newString)
		if err != nil {
			// ... handle error
			panic(err)
		}

		count += i
	}
	fmt.Println(count)
}
