package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	valueA, _ := strconv.Atoi(strings.Fields(lines[0])[4])
	valueB, _ := strconv.Atoi(strings.Fields(lines[1])[4])
	factorA := 16807
	factorB := 48271
	matched := 0

	for i := 0; i < 40000000; i++ {
		valueA = (valueA * factorA) % 2147483647
		valueB = (valueB * factorB) % 2147483647

		binA := fmt.Sprintf("%016b", valueA)
		binA = binA[len(binA) - 16:]
		binB := fmt.Sprintf("%016b", valueB)
		binB = binB[len(binB) - 16:]

		if binA == binB {
			matched++
		}
	}
	log.Printf("Matches: %v", matched)
	log.Printf("Execution time: %v", time.Since(start))
}



func part2(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	valueA, _ := strconv.Atoi(strings.Fields(lines[0])[4])
	valueB, _ := strconv.Atoi(strings.Fields(lines[1])[4])
	factorA := 16807
	factorB := 48271
	matched := 0

	genA := func(value int) int {
		value = (value * factorA) % 2147483647

		for value % 4 != 0 {
			value = (value * factorA) % 2147483647
		}

		return value
	}

	genB := func(value int) int {
		value = (value * factorB) % 2147483647

		for value % 8 != 0 {
			value = (value * factorB) % 2147483647
		}

		return value
	}

	for i := 0; i < 5000000; i++ {
		valueA = genA(valueA)
		valueB = genB(valueB)

		binA := fmt.Sprintf("%016b", valueA)
		binA = binA[len(binA) - 16:]
		binB := fmt.Sprintf("%016b", valueB)
		binB = binB[len(binB) - 16:]

		if binA == binB {
			matched++
		}
	}
	log.Printf("Matches: %v", matched)
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-15")
	part1(input)
	part2(input)
}
