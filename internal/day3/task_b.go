package day3

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

func Day3b() {
	filecontent := utils.ReadFile("3", false)
	sum := 0
	digitsNeeded := 12

	for _, line := range filecontent {
		runes := []rune(line)
		runesLen := len(runes)

		digits := make([]rune, 0, digitsNeeded)
		indexStart := 0

		for i := digitsNeeded - 1; i >= 0; i-- {
			digit, indexDigit := getHighestNumA(runes[indexStart : runesLen-i])
			digits = append(digits, digit)
			indexStart = indexStart + indexDigit + 1
		}

		joltageStr := string(digits)
		// fmt.Println("joltage: " + joltageStr)
		joltage := utils.ConvertStringToNumber(joltageStr)

		sum += joltage
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}
