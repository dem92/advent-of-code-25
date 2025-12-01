package day1

import (
	"aoc25/internal/utils"
	"fmt"
)

func Day1b() {
	position := 50
	fileContent := utils.ReadFile("1", false)
	zeroCounter := 0

	for _, line := range fileContent {
		dir := []rune(line)[0]
		clicksToTurnStr := string([]rune(line)[1:])
		clicksToTurn := utils.ConvertStringToNumber(clicksToTurnStr)
		zeroPasses := 0

		switch dir {
		case 'L':
			position, zeroPasses = turnB(position, -clicksToTurn)
		case 'R':
			position, zeroPasses = turnB(position, clicksToTurn)
		}

		zeroCounter += zeroPasses
	}

	fmt.Println("Zeros: ", zeroCounter)
}

func turnB(originalPosition, clicksToTurn int) (newPosition int, zeroPasses int) {
	possibleValues := 100
	newPosition = originalPosition + clicksToTurn

	if newPosition < 0 {
		if originalPosition != 0 {
			zeroPasses++
		}

		fullTurns := newPosition / possibleValues
		if fullTurns < 0 {
			fullTurns = fullTurns * -1
		}
		zeroPasses += fullTurns
		newPosition = newPosition % possibleValues

		if newPosition < 0 {
			newPosition = possibleValues + newPosition
		}
	} else if newPosition > possibleValues {
		zeroPasses = newPosition / possibleValues
		newPosition = newPosition % possibleValues
	} else if newPosition == possibleValues {
		zeroPasses++
		newPosition = 0
	} else if newPosition == 0 {
		zeroPasses++
	}

	return
}
