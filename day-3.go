package main

import (
	"log"
	"strconv"
	"time"
)

func part1(number int) {
	start := time.Now()
	circle, maxCircleValue := 1, 0

	for {
		valuesPerRow := circle * 2 - 1
		maxCircleValue = valuesPerRow * valuesPerRow

		if maxCircleValue > number {
			break
		}

		circle++
	}

	circleSteps := circle - 1

	for {
		left := maxCircleValue - number

		if left > circleSteps / 2 {
			maxCircleValue -= circleSteps
			continue
		}

		if left < 0 {
			left *= -1
		}

		log.Printf("Steps required: %v", circleSteps * 2 - left)
		break
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func part2(number int) {
	start := time.Now()

	type Address struct {
		x, y int
	}

	calcValue := func (x int, y int, grid map[Address]int) int {
		value := 0
		coords := [8]Address{
			{x + 1, y},
			{x + 1, y + 1},
			{x, y + 1},
			{x - 1, y + 1},
			{x - 1, y},
			{x - 1, y - 1},
			{x, y - 1},
			{x + 1, y - 1},
		}

		for i := range coords {
			if v, ok := grid[coords[i]]; ok {value += v}
		}

		return value
	}

	grid := make(map[Address]int)
	x, y, circle := 1, -1, 2
	grid[Address{0, 0}] = 1
	turn, value := 0, 0

	for {
		switch turn {
		// Up - y++
		case 0:
			if y++; y == circle - 1 {
				turn++
			}
		// Left - x--
		case 1:
			if x--; x == 1 - circle {
				turn++
			}
		// Down - y--
		case 2:
			if y--; y == 1 - circle {
				turn++
			}
		// Right - x++
		case 3:
			if x++; x == circle {
				turn = 0
				circle++
			}
		}

		value = calcValue(x, y, grid)

		if value > number {
			log.Printf("Value: %v", value)
			break
		}

		grid[Address{x, y}] = value
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-3")

	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	part1(number)
	part2(number)
}
