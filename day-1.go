package main

import (
	"log"
	"strconv"
	"time"
)

func part1(input string) {
	start := time.Now()
	lastIndex := len(input) - 1
	sum := 0

	for i := range input {
		nextI := i + 1

		if i == lastIndex {
			nextI = 0
		}

		if input[i] == input[nextI] {
			digit, err := strconv.Atoi(string(input[i]))

			if err != nil {
				log.Fatal(err)
				break
			}

			sum += digit
		}
	}

	log.Printf("Sum: %v", sum)
	log.Printf("Execution time: %v", time.Since(start))
}

func part2(input string) {
	start := time.Now()
	inputLen := len(input)
	lastIndex := inputLen - 1
	offset := len(input) / 2
	sum := 0

	for i := range input {
		nextI := i + offset

		if nextI > lastIndex {
			nextI %= inputLen
		}

		if input[i] == input[nextI] {
			digit, err := strconv.Atoi(string(input[i]))

			if err != nil {
				log.Fatal(err)
				break
			}

			sum += digit
		}
	}

	log.Printf("Sum: %v", sum)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-1")
	part1(input)
	part2(input)
}
