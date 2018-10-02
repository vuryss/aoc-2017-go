package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(lines []string) {
	start := time.Now()
	numbers := make([]int, len(lines))

	for i, line := range lines {
		if number, err := strconv.Atoi(line); err == nil {
			numbers[i] = number
		}
	}

	position, steps := 0, 0

	for position < len(numbers) {
		offset := numbers[position]
		numbers[position]++
		position += offset
		steps++
	}

	log.Printf("Steps: %v", steps)

	log.Printf("Execution time: %v", time.Since(start))
}

func part2(lines []string) {
	start := time.Now()
	numbers := make([]int, len(lines))

	for i, line := range lines {
		if number, err := strconv.Atoi(line); err == nil {
			numbers[i] = number
		}
	}

	position, steps := 0, 0

	for position < len(numbers) {
		offset := numbers[position]
		if offset > 2 {
			numbers[position]--
		} else {
			numbers[position]++
		}
		position += offset
		steps++
	}

	log.Printf("Steps: %v", steps)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-5")
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}
