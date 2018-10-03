package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	registers := make(map[string]int)
	highestValue := 0

	for i := range lines {
		parts := strings.Fields(lines[i])

		if _, exists := registers[parts[0]]; !exists {
			registers[parts[0]] = 0
		}

		checkedRegisterValue, exists := registers[parts[4]]
		compareValue, _ := strconv.Atoi(parts[6])

		if !exists {
			registers[parts[4]] = 0
		}

		shouldExecute := false

		switch parts[5] {
		case ">":
			shouldExecute = checkedRegisterValue > compareValue
		case "<":
			shouldExecute = checkedRegisterValue < compareValue
		case ">=":
			shouldExecute = checkedRegisterValue >= compareValue
		case "<=":
			shouldExecute = checkedRegisterValue <= compareValue
		case "!=":
			shouldExecute = checkedRegisterValue != compareValue
		case "==":
			shouldExecute = checkedRegisterValue == compareValue
		}

		if shouldExecute {
			modifier, _ := strconv.Atoi(parts[2])
			if parts[1] == "inc" {
				registers[parts[0]] += modifier
			} else {
				registers[parts[0]] -= modifier
			}

			if registers[parts[0]] > highestValue {
				highestValue = registers[parts[0]]
			}
		}

	}

	max := 0

	for i := range registers {
		if registers[i] > max {
			max = registers[i]
		}
	}

	log.Printf("Largest value: %v", max)
	log.Printf("Highest value during execution: %v", highestValue)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-8")
	part(input)
}
