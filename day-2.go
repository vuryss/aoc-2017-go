package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	sum := 0

	for i := range lines {
		strNums := strings.Fields(lines[i])
		min, max := int(^uint(0) >> 1), 0

		for j := range strNums {
			num, err := strconv.Atoi(strNums[j])

			if err != nil {
				log.Fatal(err)
			}

			if num < min {
				min = num
			}

			if num > max {
				max = num
			}
		}

		sum += max - min
	}

	log.Printf("Checksum: %v", sum)

	log.Printf("Execution time: %v", time.Since(start))
}

func part2(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	sum := 0

	for i := range lines {
		strNums := strings.Fields(lines[i])
		numNumbers := len(strNums)
		numbers := make([]int, len(strNums))

		for j := range strNums {
			num, err := strconv.Atoi(strNums[j])

			if err != nil {
				log.Fatal(err)
			}

			numbers[j] = num
		}

		for j := range numbers {
			for k := j + 1; k < numNumbers; k++ {
				if numbers[j] % numbers[k] == 0 {
					sum += numbers[j] / numbers[k]
					continue
				}

				if numbers[k] % numbers[j] == 0 {
					sum += numbers[k] / numbers[j]
				}
			}
		}
	}

	log.Printf("Checksum: %v", sum)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-2")
	part1(input)
	part2(input)
}
