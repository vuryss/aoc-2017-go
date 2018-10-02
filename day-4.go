package main

import (
	"log"
	"sort"
	"strings"
	"time"
)

func part1(lines []string) {
	start := time.Now()
	var mapWords map[string]bool
	sum := 0

	for i := range lines {
		words := strings.Fields(lines[i])
		mapWords = make(map[string]bool)

		for j := range words {
			mapWords[words[j]] = true
		}

		if len(mapWords) == len(words) {
			sum++
		}
	}

	log.Printf("Valid passwords: %v", sum)

	log.Printf("Execution time: %v", time.Since(start))
}

func part2(lines []string) {
	start := time.Now()
	var mapWords map[string]bool
	sum := 0

	for i := range lines {
		words := strings.Fields(lines[i])
		mapWords = make(map[string]bool)

		for j := range words {
			chars := []rune(words[j])
			sort.Slice(chars, func (i, j int) bool { return chars[i] < chars[j]})
			words[j] = string(chars)
			mapWords[words[j]] = true
		}

		if len(mapWords) == len(words) {
			sum++
		}
	}

	log.Printf("Valid passwords: %v", sum)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-4")
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}
