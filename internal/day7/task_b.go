package day7

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

type day7b struct {
	runeLines    [][]rune
	splittersHit map[string]int
}

func Day7b() {
	filecontent := utils.ReadFile("7", false)
	runeLines = utils.Get2dRuneArray(filecontent)
	d7b := day7b{
		runeLines:    runeLines,
		splittersHit: map[string]int{},
	}

	startY := 1
	startX := findStartRune(runeLines[0])

	sum := d7b.FollowBeam(startY, startX)
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func (d day7b) FollowBeam(y, x int) int {
	if x < 0 || x == len(d.runeLines[0]) {
		return 0
	}
	if y == len(d.runeLines) {
		return 1
	}

	nextY := y + 1
	if nextRune := d.runeLines[y][x]; nextRune != '^' {
		return d.FollowBeam(nextY, x)
	}

	splitterCoord := fmt.Sprintf("%d%d", y, x)
	recordedEnds := d.splittersHit[splitterCoord]

	if recordedEnds != 0 {
		return recordedEnds
	}

	ends := d.FollowBeam(nextY, x-1) + d.FollowBeam(nextY, x+1)
	d.splittersHit[splitterCoord] = ends
	return ends
}
