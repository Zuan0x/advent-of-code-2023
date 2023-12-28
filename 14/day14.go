package main

import (
	"fmt"
	"os"
	"strings"
)

type coordinates struct {
	x, y int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	rows := strings.Split(string(file), "\n")

	fmt.Println(task1(rows))
}

func task1(rows []string) int {
	new := shiftRocks(rows)
	fmt.Println(calculateWeight(new))
	return 0
}

func shiftRocks(rows []string) []string {
	newRows := rows

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			// if rows[i][j] == '#' {
			// 	//Maintain postion of square rocks
			// 	if len(newRows[i]) == 0 {
			// 		newRows[i] = strings.Join([]string{rows[i][:j], "#", rows[i][j+1:]}, "")
			// 	} else {
			// 		newRows[i] = strings.Join([]string{newRows[i][:j], "#", newRows[i][j+1:]}, "")
			// 	}
			// }
			if rows[i][j] == 'O' {
				//Replace round rock with empty space

				if i == 0 {
					newRows[0] = strings.Join([]string{newRows[0][:j], "O", newRows[0][j+1:]}, "")
				} else {
					newRows[i] = strings.Join([]string{rows[i][:j], ".", rows[i][j+1:]}, "")
				}
				//Look for new position
				for k := i - 1; k >= 0; k-- {
					if newRows[k][j] == 'O' || newRows[k][j] == '#' {
						fmt.Println("Found rock at", k, j, "checking", i)
						newRows[k+1] = strings.Join([]string{newRows[k+1][:j], "O", newRows[k+1][j+1:]}, "")
						break
					} else if k == 0 {
						fmt.Println("Found ground at", k, j, "checking", i)
						newRows[k] = strings.Join([]string{newRows[k][:j], "O", newRows[k][j+1:]}, "")
						break
					}
				}
			}
		}
	}

	for i := 0; i < len(newRows); i++ {
		fmt.Println(newRows[i])
	}

	return newRows
}

func calculateWeight(rows []string) int {
	weight := 0

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == 'O' {
				weight += len(rows[i]) - i
			}
		}
	}

	return weight
}
