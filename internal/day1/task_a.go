package day1

import (
	"aoc25/internal/utils"
	"fmt"
)

const maxValue = 99

func Day1a() {
	position := 50
	fileContent := utils.ReadFile("1", false)
	zeroCounter := 0

	for _, line := range fileContent {
		dir := []rune(line)[0]
		clicksToTurnStr := string([]rune(line)[1:])
		clicksToTurn := utils.ConvertStringToNumber(clicksToTurnStr)

		fmt.Println("clicksToTurn: ", clicksToTurn)

		switch dir {
		case 'L':
			position = turnA(position, -clicksToTurn)
		case 'R':
			position = turnA(position, clicksToTurn)
		}

		if position == 0 {
			zeroCounter++
		}

		fmt.Println("Position: ", position)
	}

	fmt.Println("Zeros: ", zeroCounter)
}

func turnA(position, clicksToTurn int) int {
	position = position + clicksToTurn

	for position > maxValue {
		position = position - 1 - maxValue
	}
	for position < 0 {
		position = position + maxValue + 1
	}

	return position
}
