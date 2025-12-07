package day7

import (
	"aoc25/internal/utils"
	"fmt"
	"strconv"
)

type day7a struct {
	runeLines    [][]rune
	splittersHit map[string]bool
}

func Day7a() {
	filecontent := utils.ReadFile("7", false)
	runeLines := utils.Get2dRuneArray(filecontent)
	d7a := day7a{
		runeLines:    runeLines,
		splittersHit: map[string]bool{},
	}

	startY := 1
	startX := findStartRune(runeLines[0])

	d7a.FollowBeam(startY, startX)

	sum := len(d7a.splittersHit)
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func (d day7a) FollowBeam(y, x int) {
	if y == len(d.runeLines) || x < 0 || x == len(d.runeLines[0]) {
		return
	}

	nextY := y + 1
	if nextRune := d.runeLines[y][x]; nextRune != '^' {
		d.FollowBeam(nextY, x)
		return
	}

	if d.registerSplitterHit(y, x) {
		d.FollowBeam(nextY, x-1)
		d.FollowBeam(nextY, x+1)
	}
}

func (d day7a) registerSplitterHit(y, x int) bool {
	splitterCoord := fmt.Sprintf("%d%d", y, x)

	registered := d.splittersHit[splitterCoord]
	if !registered {
		d.splittersHit[splitterCoord] = true
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
