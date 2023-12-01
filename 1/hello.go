package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// input := "6oneightsr"

	// re1 := regexp.MustCompile(`([a-z]+|\d)`)

	// matches := re1.FindAllStringSubmatch(input, -1)

	// fmt.Println(containsWord(matches[1][0]))

	task2b()
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

func task2() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	count := 0

	// Split the file into lines
	lines := strings.Split(string(file), "\n")

	// Define a regular expression to match the words "one" through "nine" and the digits 0 through 9
	re1 := regexp.MustCompile(`([a-z]+|\d)`)
	// using for loop
	for i := 0; i < len(lines); i++ {
		//Split the string into words and numbers
		matches := re1.FindAllStringSubmatch(lines[i], -1)

		fmt.Println("  ")

		newString := ""

		//Check if the first subarray contains words
		newString += containsNum(matches[0][0])

		//Check if the first subarray contains numbers
		newString += containsWord(matches[0][0])

		//Check if the last subarray contains words
		newString += containsNum(matches[len(matches)-1][0])

		//Check if the last subarray contains numbers
		newString += containsWord(matches[len(matches)-1][0])

		fmt.Println("NEw", lines[i], newString)

		// string to int
		i, err := strconv.Atoi(newString)
		if err != nil {
			// ... handle error
			panic(err)
		}

		count += i
	}
	println(count)
}

func containsNum(s string) string {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	fmt.Println("containsNum", s)
	if strings.Contains(string(digits), s) {
		//Append newString with only the digits
		return s
	}
	return ""
}

func containsWord(s string) string {

	wordToNumberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	//Gradually increase the max number
	println("containsWord", wordToNumberMap, s)
	for i := 0; i < len(s)-1; i++ {
		println(s[len(s)-i:])
		re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine|\d)`)
		matches := re.FindAllStringSubmatch(s[len(s)-i-1:], -1)
		println(matches)
		if wordToNumberMap[matches[0][0]] != 0 {
			println("containsWord", wordToNumberMap[matches[0][0]])
		}
	}
	return ""
}

func task2b() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Split the file into lines
	lines := strings.Split(string(file), "\n")

	//input := "6oneightsr" //"xninetwonine234nvtlzxzczx" //"6oneightsr"

	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	wordToNumberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine)`)

	count := 0

	for i := 0; i < len(lines); i++ {
		input := lines[i] //"two16shfjsixglgvkjtxkvdlqtwotwo" //
		newString := ""

		//Loop forwards through the string
		for i := 0; i < len(input); i++ {
			//println(input[:i])

			//Check if the word contains a number
			matches := re.FindStringSubmatch(input[:i])
			if len(matches) > 0 {
				newString += wordToNumberMap[matches[0]]
				break
			}

			//Check if the character is a digit
			if strings.Contains(string(digits), string(input[i])) {
				//Append newString with only the digit
				newString += string(input[i])
				break
			}
		}

		//Loop backwards through the string
		for i := 0; i < len(input); i++ {
			//println(input[len(input)-i-1:])

			//Check if the word contains a number
			matches := re.FindStringSubmatch(input[len(input)-i-1:])
			if len(matches) > 0 {
				newString += wordToNumberMap[matches[0]]
				break
			}

			//Check if the character is a digit
			if strings.Contains(string(digits), string(input[len(input)-i-1])) {
				//Append newString with only the digit
				newString += string(input[len(input)-i-1])
				break
			}
		}
		println(" ")
		println(input, newString)
		// string to int
		i, err := strconv.Atoi(newString)
		if err != nil {
			// ... handle error
			panic(err)
		}

		count += i
	}
	println(count)
}
