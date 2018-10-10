package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func findCombinations(conn int, remaining []Point2) [][]Point2 {
	combinations := make([][]Point2, 0)

	for i := range remaining {
		if remaining[i].x == conn || remaining[i].y == conn {
			combinations = append(combinations, []Point2{remaining[i]})
			list := copyList(remaining)
			list = removeFromList(list, remaining[i])
			found := make([][]Point2, 0)

			if remaining[i].x == conn {
				found = findCombinations(remaining[i].y, list)
			} else {
				found = findCombinations(remaining[i].x, list)
			}

			for j := range found {
				newList := []Point2{remaining[i]}
				newList = append(newList, found[j]...)
				combinations = append(combinations, newList)
			}
		}
	}

	return combinations
}

func removeFromList(a []Point2, conn Point2) []Point2 {
	for i := range a {
		if a[i] == conn {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}

	return a
}

func copyList(a []Point2) []Point2 {
	b := make([]Point2, len(a))
	copy(b, a)
	return b
}

func day24part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	connectors := make([]Point2, len(lines))

	for i := range lines {
		parts := strings.Split(lines[i], "/")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		connectors[i] = Point2{first, second}
	}

	combinations := findCombinations(0, connectors)
	maxLength := 0
	max := 0

	for i := range combinations {
		sum := 0
		for j := range combinations[i] {
			sum += combinations[i][j].x + combinations[i][j].y
		}

		if len(combinations[i]) > maxLength {
			maxLength = len(combinations[i])
		}

		if sum > max {
			max = sum
		}
	}

	log.Printf("Max strength: %v", max)
	max = 0

	for i := range combinations {
		if len(combinations[i]) < maxLength {
			continue
		}

		sum := 0

		for j := range combinations[i] {
			sum += combinations[i][j].x + combinations[i][j].y
		}

		if sum > max {
			max = sum
		}
	}

	log.Printf("Max strength of longest: %v", max)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-24")
	day24part1(input)
}
