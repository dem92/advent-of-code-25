package day2

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day2b() {
	filecontent := utils.ReadFile("2", false)[0]
	ranges := strings.Split(filecontent, ",")
	sum := 0

	for _, ra := range ranges {
		sum += getRangeSumB(ra)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func getRangeSumB(ra string) (sum int) {
	parts := strings.Split(ra, "-")
	startStr, endStr := parts[0], parts[1]
	start, end := utils.ConvertStringToNumber(startStr), utils.ConvertStringToNumber(endStr)
	invalidIds := map[int]int{}

	for id := start; id <= end; id++ {
		idStr := strconv.Itoa(id)
		idLen := len(idStr)
		halfLen := idLen / 2
		idRunes := []rune(idStr)

		for j := range idStr {
			if j == halfLen {
				break
			}
			if idLen%(j+1) != 0 {
				continue
			}

			pattern := string(idRunes[:j+1])
			numStr := ""
			for range idLen / len(pattern) {
				numStr += pattern
			}

			num := utils.ConvertStringToNumber(numStr)
			if num < start || num > end {
				continue
			}
			invalidIds[num] = num
		}
	}

	for k := range invalidIds {
		sum += k
	}

	return
}
