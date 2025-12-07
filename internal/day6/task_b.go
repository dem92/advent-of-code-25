package day6

import (
	"aoc25/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day6b() {
	filecontent := utils.ReadFile("6", false)
	contLen := len(filecontent)

	opsStr := filecontent[contLen-1]
	opsRegex := regexp.MustCompile(`\+|\*`)
	ops := opsRegex.FindAllString(opsStr, -1)

	sums := make([]int, 0, len(ops))
	numStrs := make([]string, len(filecontent[0]))

	for _, line := range filecontent[:contLen-1] {
		for i, r := range []rune(line) {
			numStrs[i] += string(r)
		}
	}

	numRegex := regexp.MustCompile(`\d+`)
	opIndex := 0
	newProblem := true
	for _, line := range numStrs {
		if numRegex.FindString(line) == "" {
			newProblem = true
			opIndex++
			continue
		}

		numStr := strings.TrimSpace(line)
		num := utils.ConvertStringToNumber(numStr)

		if newProblem {
			newProblem = false
			sums = append(sums, num)
			continue
		}

		sums[opIndex] = calculate(sums[opIndex], num, ops[opIndex])
	}

	sum := 0
	for _, s := range sums {
		sum += s
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}
