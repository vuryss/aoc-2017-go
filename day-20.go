package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Particle struct {
	x, y, z int
	vx, vy, vz int
	ax, ay, az int
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func day20part1(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	particles := make([]Particle, len(lines))
	re := regexp.MustCompile(`-?\d+`)

	for i := range lines {
		matches := re.FindAllString(lines[i], -1)
		matchesInt := make([]int, len(matches))

		for j := range matches {
			matchesInt[j], _ = strconv.Atoi(matches[j])
		}

		particles[i] = Particle{
			matchesInt[0], matchesInt[1], matchesInt[2],
			matchesInt[3], matchesInt[4], matchesInt[5],
			matchesInt[6], matchesInt[7], matchesInt[8],
		}
	}

	lowestSum, pIndex := 1000, 0

	for i := range particles {
		sum := abs(particles[i].ax) + abs(particles[i].ay) + abs(particles[i].az)

		if sum < lowestSum {
			lowestSum = sum
			pIndex = i
		}
	}

	log.Printf("Particle: %v", pIndex)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-20")
	day20part1(input)
}
