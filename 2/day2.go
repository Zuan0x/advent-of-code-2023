package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//task1()
	task2()
}

func task1() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	bag := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	// Split the file into lines
	lines := strings.Split(string(file), "\n")

	count := 0

	// using for loop
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		line := strings.Split(string(lines[i]), ": ")

		//Extract game number
		gameId := line[0][5:]
		fmt.Println(gameId)

		possible := true

		//Extract different pickups
		hands := strings.Split(string(line[1]), "; ")
		fmt.Println(hands)
		for j := 0; j < len(hands); j++ {
			hand := strings.Split(string(hands[j]), ", ")
			//Extract pickup number
			for k := 0; k < len(hand); k++ {
				pickup := strings.Split(string(hand[k]), " ")
				//fmt.Println(pickup[1], pickup[0], bag[pickup[1]])
				amount, err := strconv.Atoi(pickup[0])
				if err != nil {
					// ... handle error
					panic(err)
				}
				if amount > bag[pickup[1]] {
					fmt.Println("Too many", pickup[1], "in", gameId, ":", pickup[0], ">", bag[pickup[1]])
					possible = false
					break
				}
			}
		}

		if possible {
			gameIdInt, err := strconv.Atoi(gameId)
			if err != nil {
				// ... handle error
				panic(err)
			}
			count += gameIdInt
		}
	}

	println("FINAL", count)
}

func task2() {
	// Open the text file for reading
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Split the file into lines
	lines := strings.Split(string(file), "\n")

	count := 0

	// using for loop
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		line := strings.Split(string(lines[i]), ": ")

		//Extract game number
		gameId := line[0][5:]
		fmt.Println(gameId)

		//Reset the bag for every game
		bag := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		//Extract different pickups
		hands := strings.Split(string(line[1]), "; ")
		fmt.Println(hands)
		for j := 0; j < len(hands); j++ {
			hand := strings.Split(string(hands[j]), ", ")
			//Extract pickup number
			for k := 0; k < len(hand); k++ {
				pickup := strings.Split(string(hand[k]), " ")
				amount, err := strconv.Atoi(pickup[0])
				if err != nil {
					panic(err)
				}
				if amount > bag[pickup[1]] {
					bag[pickup[1]] = amount
				}
			}
		}

		//Iterate over bag to get the total number of balls
		total := 1
		for _, value := range bag {
			total *= value
		}

		count += total
	}

	println("FINAL", count)
}
