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
	cards := strings.Split(string(file), "\n")

	//ask1(cards)
	task2(cards)
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

func task2(cards []string) {
	//Create a map to keep track of the number of cards we have
	cardMap := make(map[string]int)

	for i := 0; i < len(cards); i++ {
		//Split each card into winning and your numbers
		allCards := strings.Split(cards[i], ": ")
		cardIndex := strings.TrimSpace(strings.Split(string(allCards[0]), "Card")[1])
		cardMap[cardIndex] += 1
		if cardMap[cardIndex] != 0 {

			cardNumbers := strings.Split(string(allCards[1]), " | ")
			winningNumbers := strings.Split(string(cardNumbers[0]), " ")
			yourNumbers := strings.Split(string(cardNumbers[1]), " ")

			//Create a map to keep track of the numbers
			numbers := make(map[string]int)

			//Loop through each number in the winning numbers
			for j := 0; j < len(winningNumbers); j++ {
				if string(winningNumbers[j]) != "" {
					//Add the number to the map
					numbers[strings.TrimSpace((winningNumbers[j]))] = 0
				}
			}

			//Loop through each number in your numbers
			for j := 0; j < len(yourNumbers); j++ {
				//If the number is in the map, increment the count
				if _, ok := numbers[string(yourNumbers[j])]; ok {
					numbers[strings.TrimSpace((yourNumbers[j]))]++
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

			//Add the number of cards to the map
			for i := 0; i < count; i++ {
				cardIndexInt, err := strconv.Atoi(cardIndex)
				if err != nil {
					// ... handle error
					panic(err)
				}

				newIndex := cardIndexInt + i + 1
				cardMap[strconv.Itoa(newIndex)] += cardMap[cardIndex]
			}
		}

	}

	total := 0

	for _, value := range cardMap {
		//Add the number to the total
		total += value
	}

	fmt.Println(total)

}
