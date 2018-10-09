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

func (p *Particle) tick() {
	p.vx += p.ax
	p.vy += p.ay
	p.vz += p.az
	p.x += p.vx
	p.y += p.vy
	p.z += p.vz
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func day20(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	particles := make(map[int]Particle)
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

	iterations := 100

	for i := 0; i < iterations; i++ {
		points := make(map[Point3]int)
		collided := make(map[int]bool, 0)

		for j, particle := range particles {
			particle.tick()
			particles[j] = particle
			point := Point3{particle.x, particle.y, particle.z}

			if index, exists := points[point]; exists {
				collided[j], collided[index] = true, true
				continue
			}

			points[point] = j
		}

		for j := range collided {
			delete(particles, j)
		}

		if i == iterations - 1 {
			log.Printf("After %v steps, particles left: %v", iterations, len(particles))
		}
	}

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-20")
	day20(input)
}
