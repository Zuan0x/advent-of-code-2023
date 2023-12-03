package main

import (
	"fmt"
	"os"
	"regexp"
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
	lines := strings.Split(string(file), "\n")

	task1(lines)
	//task2()
}

func isAdjacent(i, j int, lines []string) bool {
	for row := i - 1; row <= i+1; row++ {
		for col := j - 1; col <= j+1; col++ {
			if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[i]) {
				continue
			}
			if lines[row][col] != '.' && (lines[row][col] < '0' || lines[row][col] > '9') {
				return true
			}
		}
	}
	return false
}

func task1(lines []string) {
	s := 0

	//Loop through each line
	for i, line := range lines {
		re := regexp.MustCompile(`\d+`) // numbers
		matches := re.FindAllStringIndex(line, -1)

		//Loop through each number match
		for _, m := range matches {
			//Loop through each digit in the number
			for j := m[0]; j < m[len(m)-1]; j++ {
				//Check if the digit is adjacent to a symbol
				if isAdjacent(i, j, lines) {
					//Convert the number to an int and add it to the sum
					n := line[m[0]:m[len(m)-1]]
					v, _ := strconv.Atoi(n)
					s += v
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", s)
}

func getAdjacentCharacters(i, j int, rows []string) int {
	// Split the text into rows
	// Check bounds to ensure i, j are within the range
	// Check if i is within bounds

	if i < 1 || i >= len(rows) {
		return 0
	}

	// Check if j is within bounds
	if j < 1 || j >= len(rows[i-1]) {
		return 0
	}

	println(string(rows[i][j]))
	// Extract the relevant rows
	relevantRows := rows[i-1 : i+2]

	re := regexp.MustCompile(`\d+`)

	total := 0

	// Extract the adjacent characters
	for _, row := range relevantRows {
		//Check all adjacent characters expect the ones in the middle column
		fmt.Println(string(row))

		if re.MatchString(string(row[j-1])) && re.MatchString(string(row[j+1])) {
			println("BOTH", string(row[j-1]), row[j-1] > 47)

			// Find all matches in the input string
			matches := re.FindAllString(row[j-3:j+4], -1)

			for _, num := range matches {
				fmt.Println(num)
				amount, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				total += amount
			}

		} else if re.MatchString(string(row[j-1])) {
			println("LEFT", string(row[j-1]), row[j-1] > 47)
			// Find all matches in the input string
			matches := re.FindAllString(row[j-3:j+1], -1)

			for _, num := range matches {
				fmt.Println(num)
				amount, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				total += amount
			}

		} else if re.MatchString(string(row[j+1])) {
			println("RIGHT", string(row[j+1]), row[j+1] > 47)
			// Find all matches in the input string
			matches := re.FindAllString(row[j:j+4], -1)

			for _, num := range matches {
				fmt.Println(num)
				amount, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				total += amount
			}
		}
	}

	println("______________________")

	return total
}

func task2() {

}
