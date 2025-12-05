package day5

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type idrange struct {
	start uint64
	end   uint64
}

func Day5b() {
	filecontent := utils.ReadFile("5", false)
	sum := 0
	ranges := []*idrange{}

	for _, line := range filecontent {
		if line == "" {
			break
		}

		rangeParts := strings.Split(line, "-")
		start, end := strToUint(rangeParts[0]), strToUint(rangeParts[1])
		r := idrange{start: start, end: end}
		ranges = append(ranges, &r)
	}

	// Merge overlapping ranges until there are none left to merge
	for {
		oldLenRanges := len(ranges)
		ranges = mergeRanges(ranges)
		if len(ranges) == oldLenRanges {
			break
		}
	}

	for _, r := range ranges {
		sum += int(r.end - r.start + 1)
	}

	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func mergeRanges(ranges []*idrange) []*idrange {
	newRanges := []*idrange{}

	for _, oldRange := range ranges {
		start, end := oldRange.start, oldRange.end
		match := false
		for i := range newRanges {
			r := newRanges[i]
			// start < range < end
			// New range "swallows" old range, replace start and end values
			if start < r.start && end > r.end {
				r.start = start
				r.end = end
				match = true
				// start < range start <= end <= range end
				// Replace old start value with lower start value
			} else if start < r.start && end >= r.start && end <= r.end {
				r.start = start
				match = true
				// range start <= start <= range end < end
				// Replace old end value with higher end value
			} else if start >= r.start && start <= r.end && end > r.end {
				r.end = end
				match = true
				// range <= start < end <= range
				// New range is "swallowed" by exisiting range
			} else if start >= r.start && end <= r.end {
				match = true
			}
		}

		if !match {
			r := idrange{start: start, end: end}
			newRanges = append(newRanges, &r)
		}
	}

	return newRanges
}
