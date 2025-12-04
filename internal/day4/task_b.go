package day4

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

func Day4b() {
	filecontent := utils.ReadFile("4", false)
	sum := 0
	paperMap := [][]rune{}

	for _, line := range filecontent {
		paperMap = append(paperMap, []rune(line))
	}

	for {
		iterationSum := 0

		for y, yRunes := range paperMap {
			for x, r := range yRunes {
				if r != rollRune {
					continue
				}

				if canAccessRollA(x, y, paperMap) {
					paperMap[y][x] = '.'
					iterationSum++
				}
			}
		}

		if iterationSum == 0 {
			break
		}

		sum += iterationSum
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}
