package day2

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day2a() {
	filecontent := utils.ReadFile("2", false)[0]
	ranges := strings.Split(filecontent, ",")
	sum := 0

	for _, ra := range ranges {
		sum += getRangeSumA(ra)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func getRangeSumA(ra string) (sum int) {
	parts := strings.Split(ra, "-")
	startStr, endStr := parts[0], parts[1]
	start, end := utils.ConvertStringToNumber(startStr), utils.ConvertStringToNumber(endStr)

	for i := start; i <= end; i++ {
		iStr := strconv.Itoa(i)
		iLen := len(iStr)
		if iLen%2 != 0 {
			continue
		}

		halfLen := iLen / 2
		iRunes := []rune(iStr)
		leftHalf, rightHalf := iRunes[:halfLen], iRunes[halfLen:]

		if string(leftHalf) == string(rightHalf) {
			sum += i
		}
	}

	return
}
