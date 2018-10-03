package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part1(input string) {
	start := time.Now()
	parts := strings.Split(input, ",")
	list := make([]int, 256)
	position := 0
	skipSize := 0
	lengths := make([]int, len(parts))

	for i := range parts {
		length, _ := strconv.Atoi(parts[i])
		lengths[i] = length
	}

	for i := range list {
		list[i] = i
	}

	for i := range lengths {
		length := lengths[i]
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

		//log.Print(list)
	}

	log.Printf("Value: %v", list[0] * list[1])

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
	input := getInput("day-10")
	part1(input)
	//part2(input)
}
