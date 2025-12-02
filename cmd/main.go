package main

import (
	"aoc25/internal/day1"
	"aoc25/internal/day2"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Arg for task was not provided")
	}

	day := os.Args[1]
	log.Printf("Run requested for %s", day)

	switch day {
	case "1a":
		day1.Day1a()
	case "1b":
		day1.Day1b()
	case "2a":
		day2.Day2a()
	default:
		log.Fatal("Failed to match task arg with function")
	}
}
