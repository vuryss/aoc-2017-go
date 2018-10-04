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
	passed := make(map[int]bool)
	programs := make(map[int][]int)

	for i := range lines {
		parts := strings.SplitN(lines[i], " ", 3)
		strPipes := strings.Split(parts[2], ", ")
		pipes := make([]int, len(strPipes))
		program, _ := strconv.Atoi(parts[0])

		for j := range strPipes {
			pipe, _ := strconv.Atoi(strPipes[j])
			pipes[j] = pipe
		}

		programs[program] = pipes
	}

	num := 1
	passed[0] = true
	queue := []int{0}
	groupsCount := 0
	firstGroupMembers := 1

	for {
		groupsCount++

		for {
			if len(queue) == 0 {
				break
			}
			programId := queue[0]
			queue = queue[1:]
			programRel := programs[programId]

			for i := range programRel {
				if _, exists := passed[programRel[i]]; exists {
					continue
				}

				num++
				queue = append(queue, programRel[i])
				passed[programRel[i]] = true
			}
		}

		if firstGroupMembers == 1 {
			firstGroupMembers = num
		}

		if len(passed) == len(programs) {
			break
		}

		for i := range programs {
			if _, exists := passed[i]; !exists {
				queue = []int{i}
				break
			}
		}
	}

	log.Printf("Programs: %v", firstGroupMembers)
	log.Printf("Groups: %v", groupsCount)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-12")
	part1(input)
	//part2(input)
}
