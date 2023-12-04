package main

import (
	"fmt"
	"math"
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

	// Split the file into lines
	cards := strings.Split(string(file), "\n")

	task1(cards)
}

func task1(cards []string) {

	total := 0

	for i := 0; i < len(cards); i++ {
		//Split each card into winning and your numbers
		allCards := strings.Split(cards[i], ": ")
		cardNumbers := strings.Split(string(allCards[1]), " | ")
		fmt.Println(string(cardNumbers[0]), "___", cardNumbers[1])
		winningNumbers := strings.Split(string(cardNumbers[0]), " ")
		yourNumbers := strings.Split(string(cardNumbers[1]), " ")

		fmt.Println(winningNumbers, yourNumbers)

		//Create a map to keep track of the numbers
		numbers := make(map[string]int)

		//Loop through each number in the winning numbers
		for j := 0; j < len(winningNumbers); j++ {
			if string(winningNumbers[j]) != "" {
				//Add the number to the map
				numbers[string(winningNumbers[j])] = 0
			}
		}

		//Loop through each number in your numbers
		for j := 0; j < len(yourNumbers); j++ {
			//If the number is in the map, increment the count
			if _, ok := numbers[string(yourNumbers[j])]; ok {
				numbers[string(yourNumbers[j])]++
			}
		}

		count := 0
		//Loop through each number in the map
		for _, value := range numbers {
			if value != 0 {
				//Add the number to the total
				count++

			}
		}

		if count > 0 {
			total += int(math.Pow(float64(2), float64(count-1)))
		}

	}

	fmt.Println(total)
}
