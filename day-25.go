package main

import (
	"log"
	"time"
)

func day25part1(input string) {
	start := time.Now()
	state := 'A'
	steps := 12683008
	tape := make(map[int]int)
	position := 0

	for i := 0; i < steps; i++ {
		switch state {
		case 'A':
			state = 'B'
			if tape[position] == 0 {
				tape[position] = 1
				position++
				break
			}
			tape[position] = 0
			position--
		case 'B':
			if tape[position] == 0 {
				tape[position] = 1
				position--
				state = 'C'
				break
			}

			tape[position] = 0
			position++
			state = 'E'

		case 'C':
			if tape[position] == 0 {
				tape[position] = 1
				position++
				state = 'E'
				break
			}

			tape[position] = 0
			position--
			state = 'D'

		case 'D':
			if tape[position] == 0 {
				tape[position] = 1
				position--
				state = 'A'
				break
			}

			tape[position] = 1
			position--
			state = 'A'

		case 'E':
			if tape[position] == 0 {
				tape[position] = 0
				position++
				state = 'A'
				break
			}

			tape[position] = 0
			position++
			state = 'F'

		case 'F':
			if tape[position] == 0 {
				tape[position] = 1
				position++
				state = 'E'
				break
			}

			tape[position] = 1
			position++
			state = 'A'
		}
	}

	sum := 0

	for i := range tape {
		sum += tape[i]
	}

	log.Printf("Sum: %v", sum)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-25")
	day25part1(input)
}
