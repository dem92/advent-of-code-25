package day4

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

const rollRune = '@'

func Day4a() {
	filecontent := utils.ReadFile("4", false)
	sum := 0
	paperMap := [][]rune{}

	for _, line := range filecontent {
		paperMap = append(paperMap, []rune(line))
	}

	// utils.PrintCharMap(paperMap)

	for y, yRunes := range paperMap {
		for x, r := range yRunes {
			if r != rollRune {
				continue
			}

			if canAccessRollA(x, y, paperMap) {
				sum++
			}
		}
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func canAccessRollA(x, y int, paperMap [][]rune) bool {
	adjacentRolls := 0

	for _, dir := range utils.Directions {
		dirX, dirY := x+dir.X, y+dir.Y

		if dirX < 0 || dirY < 0 || dirX >= len(paperMap[0]) || dirY >= len(paperMap) {
			continue
		}

		dirRune := paperMap[dirY][dirX]

		if dirRune == rollRune {
			adjacentRolls++
		}
	}

	return adjacentRolls < 4
}
