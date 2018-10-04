package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(input string) {
	start := time.Now()

	// Parse layers [layer - range]
	lines := strings.Split(input, "\n")
	layers := make(map[int]int)
	steps := 0

	for i := range lines {
		parts := strings.Split(lines[i], ": ")
		depth, _ := strconv.Atoi(parts[0])
		rng, _ := strconv.Atoi(parts[1])
		layers[depth] = rng
		steps = depth
	}

	// Init scanners [layer - position]
	scanners := make(map[int]int)
	scannersDirection := make(map[int]bool)

	for i := range layers {
		scanners[i] = 0
		scannersDirection[i] = true
	}

	// Run
	severity := 0

	for i := 0; i <= steps; i++ {
		// Check if caught
		if scannerPosition, exists := scanners[i]; exists && scannerPosition == 0 {
			severity += i * layers[i]
		}

		// Move scanners
		for j := range scanners {
			if scannersDirection[j] {
				scanners[j]++

				if scanners[j] + 1 == layers[j] {
					scannersDirection[j] = false
				}
			} else {
				scanners[j]--

				if scanners[j] == 0 {
					scannersDirection[j] = true
				}
			}
		}
	}

	log.Printf("Severity: %v", severity)

	log.Printf("Execution time: %v", time.Since(start))
}

func part2(input string) {
	start := time.Now()

	// Parse layers [layer - range]
	lines := strings.Split(input, "\n")
	layers := make(map[int]int)
	layersIndex := make([]int, len(lines))

	for i := range lines {
		parts := strings.Split(lines[i], ": ")
		depth, _ := strconv.Atoi(parts[0])
		rng, _ := strconv.Atoi(parts[1])
		layers[depth] = rng
		layersIndex[i] = depth
	}

	startSeconds := 0

	CHECKPOINT:
	for {
		for _, i := range layersIndex {
			check := startSeconds + i
			rangeInterval := (layers[i] - 1) * 2
			if check % rangeInterval == 0 {
				startSeconds++
				continue CHECKPOINT
			}
		}

		log.Printf("Passed: %v", startSeconds)
		break
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-13")
	part1(input)
	part2(input)
}
