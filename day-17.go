package main

import (
	"log"
	"strconv"
	"time"
)

func day17part1(input string) {
	start := time.Now()
	steps, _ := strconv.Atoi(input)
	buffer := []int{0}
	position := 0

	for i := 1; i <= 2017; i++ {
		itemInBuffer := len(buffer)
		position = (position + steps) % itemInBuffer

		buffer = append(buffer, 0)
		copy(buffer[position+2:], buffer[position + 1:])
		buffer[position + 1] = i

		position++
	}

	for i := range buffer {
		if buffer[i] == 2017 {
			log.Printf("Value %v", buffer[i+1])
			break
		}
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func day17part2(input string) {
	start := time.Now()
	steps, _ := strconv.Atoi(input)
	itemInBuffer := 1
	position := 0
	afterZero := 0

	for i := 1; i <= 50000000; i++ {
		position = (position + steps) % itemInBuffer

		if position == 0 {
			afterZero = i
		}

		itemInBuffer++
		position++
	}

	log.Printf("Value after 0: %v", afterZero)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-17")
	day17part1(input)
	day17part2(input)
}
