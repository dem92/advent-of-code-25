package day3

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

func Day3a() {
	filecontent := utils.ReadFile("3", false)
	sum := 0

	for _, line := range filecontent {
		runes := []rune(line)
		runesLen := len(runes)

		first, indexFirst := getHighestNumA(runes[:runesLen-1])
		second, _ := getHighestNumA(runes[indexFirst+1:])

		joltageStr := string([]rune{first, second})
		// fmt.Println("joltage: " + joltageStr)
		joltage := utils.ConvertStringToNumber(joltageStr)

		sum += joltage
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func getHighestNumA(nums []rune) (num rune, index int) {
	highestNum := '0'
	indexHighestNum := 0

	for i, r := range nums {
		if r > highestNum {
			highestNum = r
			indexHighestNum = i
		}
	}

	return highestNum, indexHighestNum
}
