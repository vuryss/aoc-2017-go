package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func hash(input string) string {
	lengths := []rune(input)
	lengths = append(lengths, 17, 31, 73, 47, 23)
	list := make([]int, 256)
	position := 0
	skipSize := 0

	for i := range list {
		list[i] = i
	}

	for round := 0; round < 64; round++ {
		for i := range lengths {
			length := int(lengths[i])
			seq := make([]int, 0)

			// To end of the list
			end := position + length - 1
			wrap := false

			if end > 255 {
				wrap = true
				end = 255
			}

			for i := position; i <= end; i++ {
				seq = append(seq, list[i])
			}

			// From the beginning
			if wrap {
				end = position + length - 256

				for i := 0; i <= end; i++ {
					seq = append(seq, list[i])
				}
			}

			// Reverse sequence
			halfLength := length / 2
			for i := 0; i < halfLength; i++ {
				seq[i], seq[length - 1 - i] = seq[length - 1 - i], seq[i]
			}

			// Save to list
			if wrap {
				for i := position; i < 256; i++ {
					list[i], seq = seq[0], seq[1:]
				}

				for i := 0; i <= end; i++ {
					list[i], seq = seq[0], seq[1:]
				}
			} else {
				for i := position; i <= end; i++ {
					list[i], seq = seq[0], seq[1:]
				}
			}

			position += length + skipSize
			position %= 256
			skipSize++
		}
	}

	denseHash := make([]int, 16)
	hash := ""

	for i := 0; i < 16; i++ {
		for j := i * 16; j < (i + 1) * 16; j++ {
			denseHash[i] ^= list[j]
		}
	}

	for i := range denseHash {
		hash += fmt.Sprintf("%02x", denseHash[i])
	}

	return hash
}

type Point struct {
	x, y int
}

func part(input string) {
	start := time.Now()

	hexTable := make(map[rune]string)

	for i := 0; i < 16; i++ {
		hexTable[[]rune(strconv.FormatInt(int64(i), 16))[0]] = fmt.Sprintf("%04b", i)
	}

	grid := [128][128]bool{}
	count := 0

	for i := 0; i < 128; i++ {
		hashString := hash(input + "-" + strconv.Itoa(i))

		for j := range hashString {
			binaryString := []rune(hexTable[rune(hashString[j])])

			for k := range binaryString {
				grid[i][j*4+k] = binaryString[k] == '1'

				if grid[i][j*4+k] {
					count++
				}
			}
		}
	}

	used := map[Point]bool{}
	groups := 0

	var addAdjacent func(x, y int)

	addAdjacent = func(x, y int) {
		point := Point{x - 1, y}

		if _, exists := used[point]; !exists && x > 0 && grid[point.x][point.y] {
			used[point] = true
			addAdjacent(point.x, point.y)
		}

		point = Point{x + 1, y}

		if _, exists := used[point]; !exists && x < 127 && grid[point.x][point.y] {
			used[point] = true
			addAdjacent(point.x, point.y)
		}

		point = Point{x, y + 1}

		if _, exists := used[point]; !exists && y < 127 && grid[point.x][point.y] {
			used[point] = true
			addAdjacent(point.x, point.y)
		}

		point = Point{x, y - 1}

		if _, exists := used[point]; !exists && y > 0 && grid[point.x][point.y] {
			used[point] = true
			addAdjacent(point.x, point.y)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				point := Point{i, j}

				if _, exists := used[point]; exists {
					continue
				}

				groups++
				used[point] = true
				addAdjacent(i, j)
			}
		}
	}

	log.Printf("Count: %v", count)
	log.Printf("Groups: %v", groups)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-14")
	part(input)
}
