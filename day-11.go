package main

import (
	"log"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Node struct {
	point Point
	distance int
}

func part1(input string) {
	start := time.Now()
	directions := strings.Split(input, ",")
	x, y := 0, 0
	longestDistance, lastDistance := 0, 0

	for i := range directions {
		dir := []rune(directions[i])

		//log.Printf("Coords: %v %v", x, y)

		distance := 0
		node := Node{Point{0, 0}, distance}
		queue := []Node{node}
		point := Point{0, 0}
		added := make(map[Point]bool)
		added[point] = true

		for {
			node = queue[0]
			queue = queue[1:]

			if node.point.x == x && node.point.y == y {
				lastDistance = node.distance
				if lastDistance > longestDistance {
					longestDistance = lastDistance
				}
				break
			}

			distance = node.distance + 1

			if x - node.point.x == 0 && y - node.point.y > 0 {
				// N
				point = Point{node.point.x, node.point.y + 2}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}

			if x - node.point.x > 0 && y - node.point.y >= 0 {
				// NE
				point = Point{node.point.x + 1, node.point.y + 1}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}

			if x - node.point.x > 0 && y - node.point.y <= 0 {
				// SE
				point = Point{node.point.x + 1, node.point.y - 1}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}

			if x - node.point.x == 0 && y - node.point.y < 0 {
				// S
				point = Point{node.point.x, node.point.y - 2}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}

			if x - node.point.x < 0 && y - node.point.y <= 0 {
				// SW
				point = Point{node.point.x - 1, node.point.y - 1}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}

			if x - node.point.x < 0 && y - node.point.y >= 0 {
				// NW
				point = Point{node.point.x - 1, node.point.y + 1}
				if _, exists := added[point]; !exists {
					queue = append(queue, Node{point, distance})
					added[point] = true
				}
			}
		}

		if directions[i] == "n" {
			y += 2
			continue
		} else if directions[i] == "s" {
			y -= 2
			continue
		}

		if dir[0] == 'n' {
			y++
		} else {
			y--
		}

		if dir[1] == 'e' {
			x++
		} else {
			x--
		}
	}

	log.Printf("Distance: %v", lastDistance)
	log.Printf("Longest distance: %v", longestDistance)
	log.Printf("Execution time: %v", time.Since(start))
}

//func part2(input string) {
//	start := time.Now()
//	inputLen := len(input)
//	lastIndex := inputLen - 1
//	offset := len(input) / 2
//	sum := 0
//
//	for i := range input {
//		nextI := i + offset
//
//		if nextI > lastIndex {
//			nextI %= inputLen
//		}
//
//		if input[i] == input[nextI] {
//			digit, err := strconv.Atoi(string(input[i]))
//
//			if err != nil {
//				log.Fatal(err)
//				break
//			}
//
//			sum += digit
//		}
//	}
//
//	log.Printf("Sum: %v", sum)
//	log.Printf("Execution time: %v", time.Since(start))
//}

func main() {
	input := getInput("day-11")
	part1(input)
	//part2(input)
}
