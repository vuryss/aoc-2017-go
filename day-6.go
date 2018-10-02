package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(numbers []int) {
	start := time.Now()
	states := make(map[string]bool)
	genKey := func(numbers []int) string {
		values := make([]string, len(numbers))

		for i, v := range numbers {
			values[i] = strconv.Itoa(v)
		}

		return strings.Join(values, ".")
	}

	cycles, loopCycles := 0, 0
	foundState := ""


	for {
		key := genKey(numbers)

		if foundState != "" {
			if key == foundState {
				break
			}
			loopCycles++
		} else if _, found := states[key]; found {
			foundState = key
			loopCycles = 1
		}

		states[key] = true
		max := 0
		index := 0

		for i, v := range numbers {
			if v > max {
				max = v
				index = i
			}
		}

		value := numbers[index]
		numbers[index] = 0

		for value > 0 {
			index++
			if index == len(numbers) {
				index = 0
			}
			numbers[index]++
			value--
		}

		cycles++
	}
	log.Printf("Cycles: %v", cycles)
	log.Printf("Loop cycles: %v", loopCycles)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-6")
	strNumber := strings.Fields(input)
	numbers := make([]int, len(strNumber))

	for i, v := range strNumber {
		if number, err := strconv.Atoi(v); err == nil {
			numbers[i] = number
		}
	}

	part1(numbers)
}
