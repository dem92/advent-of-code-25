package day7

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

var runeLines [][]rune
var splittersHit = map[string]bool{}

func Day7a() {
	filecontent := utils.ReadFile("7", false)
	runeLines = utils.Get2dRuneArray(filecontent)

	startY := 1
	startX := findStartRune(runeLines[0])

	followBeam(startY, startX)

	sum := len(splittersHit)
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func followBeam(y, x int) {
	if y == len(runeLines) || x < 0 || x == len(runeLines[0]) {
		return
	}

	nextY := y + 1
	if nextRune := runeLines[y][x]; nextRune != '^' {
		followBeam(nextY, x)
		return
	}

	if registerSplitterHit(y, x) {
		followBeam(nextY, x-1)
		followBeam(nextY, x+1)
	}
}

func registerSplitterHit(y, x int) bool {
	splitterCoord := fmt.Sprintf("%d%d", y, x)

	registered := splittersHit[splitterCoord]
	if !registered {
		splittersHit[splitterCoord] = true
	}

	return !registered
}

func findStartRune(runes []rune) int {
	for i, r := range runes {
		if r == 'S' {
			return i
		}
	}

	panic("failed to find start rune")
}
