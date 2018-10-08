package main

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func day18part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	registers := make(map[rune]int)
	position, lastFrequency := 0, 0

	EXEC:
	for {
		parts := strings.Fields(lines[position])
		register := []rune(parts[1])[0]

		switch parts[0] {
		case "snd":
			lastFrequency, _ = registers[register]
		case "set":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registers[register] = intValue
		case "add":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue + intValue
		case "mul":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue * intValue
		case "mod":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue % intValue
		case "rcv":
			registerValue, _ := registers[register]

			if registerValue > 0 {
				log.Printf("Recovered value: %v", lastFrequency)
				break EXEC
			}
		case "jgz":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]

			if registerValue <= 0 {
				position++
				break
			}

			position += intValue
		}

		if parts[0] != "jgz" {
			position++
		}
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func sound(instructions []string, id int, selfQueue chan int, otherQueue chan int, wg sync.WaitGroup) {
	position := 0
	registers := make(map[rune]int)
	registers['p'] = id
	sentCount := 0

	for {
		parts := strings.Fields(instructions[position])
		register := []rune(parts[1])[0]

		switch parts[0] {
		case "snd":
			value, _ := registers[register]
			otherQueue <- value
			sentCount++

			if id == 1 {
				log.Printf("Program 1 sent %v times", sentCount)
			}
		case "set":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registers[register] = intValue
		case "add":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue + intValue
		case "mul":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue * intValue
		case "mod":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			registerValue, _ := registers[register]
			registers[register] = registerValue % intValue
		case "rcv":
			registers[register] = <- selfQueue
		case "jgz":
			intValue, error := strconv.Atoi(parts[2])

			if error != nil {
				intValue, _ = registers[[]rune(parts[2])[0]]
			}

			if register == '1' {
				position += intValue
				break
			}

			registerValue, _ := registers[register]

			if registerValue <= 0 {
				position++
				break
			}

			position += intValue
		}

		if parts[0] != "jgz" {
			position++
		}
	}

	log.Printf("Channel 1 sent %v times", sentCount)

	defer wg.Done()
}

func day18part2(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	var wg sync.WaitGroup

	channel1 := make(chan int, 100000)
	channel2 := make(chan int, 100000)
	wg.Add(2)
	go sound(lines, 0, channel1, channel2, wg)
	go sound(lines, 1, channel2, channel1, wg)
	wg.Wait()

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-18")
	day18part1(input)
	day18part2(input)
}
