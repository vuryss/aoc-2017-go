package main

import (
	"log"
	"time"
)

func part(input string) {
	start := time.Now()
	characters := []rune(input)
	nestingLevel, score, garbage := 0, 0, 0
	inGarbage, skipCharacter := false, false

	for i := range characters {
		if skipCharacter {
			skipCharacter = false
			continue
		}
		if inGarbage && characters[i] != '>' && characters[i] != '!' {
			garbage++
			continue
		}
		switch characters[i] {
		case '{':
			nestingLevel++
			score += nestingLevel
		case '}':
			nestingLevel--
		case '<':
			inGarbage = true
		case '>':
			inGarbage = false
		case '!':
			skipCharacter = true
		}
	}

	log.Printf("Score: %v", score)
	log.Printf("Garbage: %v", garbage)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-9")
	part(input)
}
