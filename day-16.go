package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func part(input string) {
	start := time.Now()
	instr := strings.Split(input, ",")
	programs := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	tempStr := ""
	memory := make(map[string][]rune)

	for j := 0; j < 1000000000; j++ {
		tempStr = string(programs)

		if j % 1000000 == 0 {
			log.Print(j)
		}

		if result, exists := memory[tempStr]; exists {
			programs = result
			continue
		}

		for i := range instr {
			item := []rune(instr[i])
			command := item[0]
			item = item[1:]
			parts := strings.Split(string(item), "/")

			switch command {
			case 's':
				x, _ := strconv.Atoi(parts[0])
				temp := append(programs[16-x:], programs[0:16-x]...)
				programs = temp
			case 'x':
				x, _ := strconv.Atoi(parts[0])
				y, _ := strconv.Atoi(parts[1])
				programs[x], programs[y] = programs[y], programs[x]
			case 'p':
				x := []rune(parts[0])[0]
				y := []rune(parts[1])[0]
				xPos, yPos := 0, 0
				for i := range programs {
					if programs[i] == x {
						xPos = i
					} else if programs[i] == y {
						yPos = i
					}
				}
				programs[xPos], programs[yPos] = programs[yPos], programs[xPos]
			}
		}

		memory[tempStr] = []rune(string(programs))

		if j == 0 {
			log.Printf("String: %v", string(programs))
		}
	}

	log.Printf("String: %v", string(programs))

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-16")
	part(input)
}
