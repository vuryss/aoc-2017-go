package main

import (
	"log"
	"strings"
	"time"
)

func day19(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)

	for i := range lines {
		line := make([]rune, 0)
		for _, char := range lines[i] {
			line = append(line, char)
		}
		grid = append(grid, line)
	}

	x, y, dir := 0, 0, 2

	for i := range grid[0] {
		if grid[0][i] == '|' {
			y = i
			break
		}
	}

	log.Printf("Starting at line %v column %v", x, y)
	letters := ""
	steps := 1

	RUN:
	for {
		char := grid[x][y]
		if char == ' ' {
			steps--
			break RUN
		} else if char != '|' && char != '-' && char != '+' {
			letters += string(char)
		}

		steps++

		switch dir {
		case 0: x--
		case 1: y--
		case 2: x++
		case 3: y++
		}

		if grid[x][y] == '+' {
			if dir != 2 && grid[x - 1][y] != ' ' {
				dir = 0
			} else if dir != 3 && grid[x][y - 1] != ' ' {
				dir = 1
			} else if dir != 0 && grid[x + 1][y] != ' ' {
				dir = 2
			} else if dir != 1 && grid[x][y + 1] != ' ' {
				dir = 3
			}
		}
	}

	log.Printf("Characters: %v", letters)
	log.Printf("Steps: %v", steps)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-19")
	day19(input)
}
