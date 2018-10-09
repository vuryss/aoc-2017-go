package main

import (
	"log"
	"strings"
	"time"
)

func day22part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	grid := make(map[Point2]bool)

	for i := range lines {
		items := []rune(lines[i])

		for j := range items {
			grid[Point2{i, j}] = items[j] == '#'
		}
	}

	currentPosition := Point2{len(lines) / 2, len(lines) / 2}
	direction := 0
	count := 0

	for i := 0; i < 10000; i++ {
		sector, exists := grid[currentPosition]

		if !exists {
			grid[currentPosition], sector = false, false
		}

		if sector {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
			count++
		}

		grid[currentPosition] = !grid[currentPosition]

		switch direction {
		case 0: currentPosition = Point2{currentPosition.x - 1, currentPosition.y}
		case 1: currentPosition = Point2{currentPosition.x, currentPosition.y + 1}
		case 2: currentPosition = Point2{currentPosition.x + 1, currentPosition.y}
		case 3: currentPosition = Point2{currentPosition.x, currentPosition.y - 1}
		}
	}

	log.Printf("Count: %v", count)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-22")
	day22part1(input)
}
