package main

import (
	"fmt"
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
	rows := strings.Split(string(file), "\n")

	task1(rows)
	//task2(string(file))
}

func task1(rows []string) {

	count := 0

	for i := 0; i < len(rows); i++ {
		//Find Start tile
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == 'S' {
				fmt.Println("Start", i, j)

				prevI := -1
				prevJ := -1
				nextI := i
				nextJ := j

				//Look for next character
				for {
					nextI, nextJ, prevI, prevJ = findNextTile(rows, nextI, nextJ, prevI, prevJ)
					count++
					if nextI == i && nextJ == j {
						break
					}
				}

				fmt.Println("Next tile", string(rows[nextI][nextJ]), nextI, nextJ)
				break
			}
		}
	}
	fmt.Println(count / 2)
}

func findNextTile(rows []string, i int, j int, prevI int, prevJ int) (int, int, int, int) {
	tilesToCheck := make([][]int, 0)

	// Make pipe type map
	//Map string to int array
	pipeMap := make(map[string][]int)

	//up, right, down, left
	pipeMap["S"] = []int{1, 1, 1, 1}
	pipeMap["F"] = []int{0, 1, 1, 0}
	pipeMap["7"] = []int{0, 0, 1, 1}
	pipeMap["L"] = []int{1, 1, 0, 0}
	pipeMap["J"] = []int{1, 0, 0, 1}
	pipeMap["-"] = []int{0, 1, 0, 1}
	pipeMap["|"] = []int{1, 0, 1, 0}

	tileType := string(rows[i][j])
	checkTiles := pipeMap[tileType]

	//Check the relevant directions based on pipe type
	if checkTiles[0] == 1 && i-1 >= 0 {
		//Only add if the pipe goes into the tile
		if pipeMap[string(rows[i-1][j])][2] == 1 {
			tilesToCheck = append(tilesToCheck, []int{i - 1, j})
		}
	}
	if checkTiles[1] == 1 {
		if pipeMap[string(rows[i][j+1])][3] == 1 {
			tilesToCheck = append(tilesToCheck, []int{i, j + 1})
		}
	}
	if checkTiles[2] == 1 {
		if pipeMap[string(rows[i+1][j])][0] == 1 {
			tilesToCheck = append(tilesToCheck, []int{i + 1, j})
		}
	}
	if checkTiles[3] == 1 && j-1 >= 0 {
		if pipeMap[string(rows[i][j-1])][1] == 1 {
			tilesToCheck = append(tilesToCheck, []int{i, j - 1})
		}
	}

	//Check if next tile is a pipe
	for _, tile := range tilesToCheck {
		//Make sure its not the previous tile
		if tile[0] != prevI || tile[1] != prevJ {
			//Check if its a pipe
			if rows[tile[0]][tile[1]] != '.' {
				fmt.Println("Found pipe", string(rows[tile[0]][tile[1]]))
				return tile[0], tile[1], i, j
			}
		}
	}
	fmt.Println("No pipe found")
	return -1, -1, -1, -1
}

func task2(file []string) {
}
