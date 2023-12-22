package main

import (
	"fmt"
	"os"
	"strings"
	/*"math"
	"path/filepath"

	"advent2023/pkg/file"*/)

type coordinates struct {
	x, y int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	patterns := strings.Split(string(file), "\n\n")

	fmt.Println(task1(patterns))
}

func task1(patterns []string) int {

	total := 0

	for _, pattern := range patterns {
		patternScore := 0
		rows := strings.Split(pattern, "\n")
		fmt.Println(rows)
		//Look for vertical symmetry line
		for i := 0; i < len(rows[0]); i++ {
			isSymmetrical := true
			for j := 0; j < len(rows); j++ {
				// fmt.Println(i, j)
				symmetricalRow := findRowSymmetry(rows[j], i, 0)
				if !symmetricalRow {
					isSymmetrical = false
					break
				}
				// fmt.Println("RESULT", symmetricalRow, rows[j][0:i+1], rows[j][i+1:])
			}
			fmt.Println("char ", i, isSymmetrical)
			if isSymmetrical {
				patternScore += (i + 1)
				fmt.Println(i + 1)
				break
			}
		}

		if patternScore == 0 {
			//Look for horizontal symmetry line
			for i := 0; i < len(rows); i++ {
				allSymmetrical := true
				for j := 0; j < len(rows[i]); j++ {
					symmetricalCol := findColSymmetry(rows, i, j, 0)
					if !symmetricalCol {
						allSymmetrical = false
						break
					}
				}
				fmt.Println("row", i, allSymmetrical)
				if allSymmetrical {
					patternScore += (i + 1) * 100
					fmt.Println((i + 1) * 100)
					break
				}
			}
		}

		total += patternScore
	}

	return total
}

func findColSymmetry(rows []string, i int, j int, depth int) bool {
	//Recursively work check from string start point whether corresponding characters are the same
	isSymmetrical := false
	//Final character cannot be symmetrical
	if depth == 0 && i == len(rows)-1 {
		return false
	}
	//If we've reached the end of the string, return true
	if i+depth+1 == len(rows) || i-depth < 0 {
		return true
	}

	//If the characters are the same, check the next pair
	if rows[i-depth][j] == rows[i+depth+1][j] {

		// fmt.Println(start-depth, string(row[start-depth]), start+depth+1, string(row[start+depth+1]))
		isSymmetrical = findColSymmetry(rows, i, j, depth+1)
	} else {
		return false
	}
	return isSymmetrical
}

func findRowSymmetry(row string, start int, depth int) bool {
	//Recursively work check from string start point whether corresponding characters are the same
	isSymmetrical := false
	//Final character cannot be symmetrical
	if depth == 0 && start == len(row)-1 {
		return false
	}
	//If we've reached the end of the string, return true
	if start+depth+1 == len(row) || start-depth < 0 {
		return true
	}

	//If the characters are the same, check the next pair
	if row[start-depth] == row[start+depth+1] {

		// fmt.Println(start-depth, string(row[start-depth]), start+depth+1, string(row[start+depth+1]))
		isSymmetrical = findRowSymmetry(row, start, depth+1)
	} else {
		return false
	}
	return isSymmetrical
}
