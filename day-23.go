package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func day23part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	registers := make(map[rune]int)
	position := 0
	count := 0

	for {
		parts := strings.Fields(lines[position])
		register := []rune(parts[1])[0]

		switch parts[0] {
		case "set":
			intValue, err := strconv.Atoi(parts[2])

			if err != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registers[register] = intValue
		case "sub":
			intValue, err := strconv.Atoi(parts[2])

			if err != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue - intValue
		case "mul":
			intValue, err := strconv.Atoi(parts[2])

			if err != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue * intValue
			count++
		case "jnz":
			intValue, err := strconv.Atoi(parts[2])

			if err != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			if register == '1' {
				position += intValue
				break
			}

			registerValue, _ := registers[register]

			if registerValue == 0 {
				position++
				break
			}

			position += intValue
		}

		if parts[0] != "jnz" {
			position++
		}

		if position >= len(lines) {
			break
		}
	}

	log.Printf("Count: %v", count)
	log.Printf("Execution time: %v", time.Since(start))
}


func main() {
	input := getInput("day-23")
	day23part1(input)
}
