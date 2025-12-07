package day6

import (
	"aoc25/internal/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Day6a() {
	filecontent := utils.ReadFile("6", false)
	contLen := len(filecontent)

	opsStr := filecontent[contLen-1]
	opsRegex := regexp.MustCompile(`\+|\*`)
	ops := opsRegex.FindAllString(opsStr, -1)

	sums := make([]int, 0, len(ops))
	numRegex := regexp.MustCompile(`\d+`)

	for i, line := range filecontent[:contLen-1] {
		numStrs := numRegex.FindAllString(line, -1)
		for j, numStr := range numStrs {
			num := utils.ConvertStringToNumber(numStr)

			if i == 0 {
				sums = append(sums, num)
			} else {
				sums[j] = calculate(sums[j], num, ops[j])
			}
		}
	}

	sum := 0
	for _, s := range sums {
		sum += s
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func calculate(a, b int, op string) int {
	if op == "+" {
		return a + b
	}
	return a * b
}
