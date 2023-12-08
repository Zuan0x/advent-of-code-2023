package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
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
	//task2(rows)
}

type Hand struct {
	cards string
	bid   int
	rank  int
}

var ranks = [7][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 2},
	{1, 2, 2},
	{1, 1, 3},
	{2, 3},
	{1, 4},
	{5},
}

func task1(lines []string) {

	var cardsByStrength = make(map[int][]Hand)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		cards := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])
		strength := calcStrength(cards, nil, false, 0)

		cardsByStrength[strength] = append(cardsByStrength[strength], Hand{cards, bid, 0})
	}

	for _, hands := range cardsByStrength {
		if len(hands) > 1 {
			sort.Slice(hands, func(i, j int) bool {
				return compareCards(hands[i].cards, hands[j].cards, getCardStrength)
			})
		}
	}

	var strengths []int
	for k := range cardsByStrength {
		strengths = append(strengths, k)
	}

	slices.Sort(strengths)

	var sum int
	rank := 1
	for _, s := range strengths {
		for _, hand := range cardsByStrength[s] {
			sum += hand.bid * rank
			rank++
		}
	}

	fmt.Println(sum)

}

func task2(rows []string) {

}

func convertStrToInt(str string) int {
	output, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return output
}

func compareCards(cardsOne, cardsTwo string, strengthFn func(c byte) int) bool {
	for i := 0; i < len(cardsOne); i++ {
		if strengthFn(cardsOne[i]) == strengthFn(cardsTwo[i]) {
			continue
		}

		return strengthFn(cardsOne[i]) < strengthFn(cardsTwo[i])
	}

	return false
}

func getCardStrength(card byte) int {
	return strings.Index("23456789TJQKA", string(card))
}

func calcStrength(cards string, numOfPairs []int, withJokers bool, jokers int) int {
	if len(cards) == 0 {
		slices.Sort(numOfPairs)
		for i, rank := range ranks {
			if slices.Compare(rank, numOfPairs) == 0 {
				return i + 1
			}
		}
	}

	matches := 1
	compared := string(cards[0])
	if compared == "J" {
		jokers++
	}
	cards = cards[1:]

	for strings.Index(cards, compared) != -1 {
		matches++
		if compared == "J" {
			jokers++
		}
		cards = strings.Replace(cards, compared, "", 1)
	}
	numOfPairs = append(numOfPairs, matches)

	return calcStrength(cards, numOfPairs, withJokers, jokers)
}
