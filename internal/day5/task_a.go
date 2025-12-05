package day5

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day5a() {
	filecontent := utils.ReadFile("5", false)
	sum := 0
	secondHalf := false
	ranges := map[uint64][]uint64{}

	for _, line := range filecontent {
		if secondHalf {
			id := strToUint(line)
			for start, ends := range ranges {
				found := false
				if id >= start {
					for _, end := range ends {
						if id <= end {
							found = true
							break
						}
					}
				}

				if found {
					sum++
					break
				}
			}
			continue
		}

		if line == "" {
			secondHalf = true
			continue
		}

		rangeParts := strings.Split(line, "-")
		start, end := strToUint(rangeParts[0]), strToUint(rangeParts[1])
		ranges[start] = append(ranges[start], end)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func strToUint(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return n
}
