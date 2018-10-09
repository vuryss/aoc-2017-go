package main

import (
	"log"
	"strings"
	"time"
)

type TwoState struct {
	x1, x2 int
	x3, x4 int
}

type ThreeState struct {
	x1, x2, x3 int
	x4, x5, x6 int
	x7, x8, x9 int
}

func flip(a TwoState) TwoState {
	a.x1, a.x2, a.x3, a.x4 = a.x2, a.x1, a.x4, a.x3
	return a
}

func flip3(a ThreeState) ThreeState {
	a.x1, a.x3, a.x4, a.x6, a.x7, a.x9 = a.x3, a.x1, a.x6, a.x4, a.x9, a.x7
	return a
}

func day21(input string) {
	start := time.Now()
	lines := strings.Split(input, "\n")
	trans2 := make(map[TwoState][]rune)
	trans3 := make(map[ThreeState][]rune)

	for i := range lines {
		lines[i] = strings.Replace(lines[i], "#", "1", -1)
		lines[i] = strings.Replace(lines[i], ".", "0", -1)
		parts := strings.Split(lines[i], " => ")

		count := strings.Split(parts[0], "/")

		from := []rune(strings.Replace(parts[0], "/", "", -1))
		target := []rune(strings.Replace(parts[1], "/", "", -1))

		if len(count) == 2 {
			state := TwoState{
				int(from[0] - '0'), int(from[1] - '0'),
				int(from[2] - '0'), int(from[3] - '0'),
			}

			trans2[state] = target

			// Flip
			state = flip(state)
			trans2[state] = target

			// Rotate once
			state = TwoState{
				int(from[2] - '0'), int(from[0] - '0'),
				int(from[3] - '0'), int(from[1] - '0'),
			}

			trans2[state] = target

			// Flip
			state = flip(state)
			trans2[state] = target

			// Rotate 2 times
			state = TwoState{
				int(from[3] - '0'), int(from[2] - '0'),
				int(from[1] - '0'), int(from[0] - '0'),
			}

			trans2[state] = target

			// Flip
			state = flip(state)
			trans2[state] = target

			// Rotate 3 times
			state = TwoState{
				int(from[1] - '0'), int(from[3] - '0'),
				int(from[0] - '0'), int(from[2] - '0'),
			}

			trans2[state] = target

			// Flip
			state = flip(state)
			trans2[state] = target

		} else {
			// Original
			state := ThreeState{
				int(from[0] - '0'), int(from[1] - '0'), int(from[2] - '0'),
				int(from[3] - '0'), int(from[4] - '0'), int(from[5] - '0'),
				int(from[6] - '0'), int(from[7] - '0'), int(from[8] - '0'),
			}
			trans3[state] = target

			// Flip
			state = flip3(state)
			trans3[state] = target

			// Rotate 1
			state = ThreeState{
				int(from[6] - '0'), int(from[3] - '0'), int(from[0] - '0'),
				int(from[7] - '0'), int(from[4] - '0'), int(from[1] - '0'),
				int(from[8] - '0'), int(from[5] - '0'), int(from[2] - '0'),
			}
			trans3[state] = target

			// Flip
			state = flip3(state)
			trans3[state] = target

			// Rotate 2
			state = ThreeState{
				int(from[8] - '0'), int(from[7] - '0'), int(from[6] - '0'),
				int(from[5] - '0'), int(from[4] - '0'), int(from[3] - '0'),
				int(from[2] - '0'), int(from[1] - '0'), int(from[0] - '0'),
			}
			trans3[state] = target

			// Flip
			state = flip3(state)
			trans3[state] = target

			// Rotate 3
			state = ThreeState{
				int(from[2] - '0'), int(from[5] - '0'), int(from[8] - '0'),
				int(from[1] - '0'), int(from[4] - '0'), int(from[7] - '0'),
				int(from[0] - '0'), int(from[3] - '0'), int(from[6] - '0'),
			}
			trans3[state] = target

			// Flip
			state = flip3(state)
			trans3[state] = target
		}
	}

	startPattern := []rune("010001111")

	iterator := func(pattern []rune, wallSize int, times int) string {
		for i := 0; i < times; i++ {
			var newPattern []rune

			if wallSize % 2 == 0 {
				regionCount := wallSize / 2
				nWallSize := regionCount * 3
				newPattern = make([]rune, nWallSize * nWallSize)

				for j := 0; j < regionCount; j++ {
					for k := 0; k < regionCount; k++ {
						a, b := j * wallSize * 2, k * 2
						stage := TwoState{
							int(pattern[a + b + 0] - '0'), int(pattern[a + b + 1] - '0'),
							int(pattern[a + b + 0 + wallSize] - '0'), int(pattern[a + b + 1 + wallSize] - '0'),
						}
						result, _ := trans2[stage]

						mod := j * nWallSize * 3 + k * 3

						for q := 0; q < 3; q++ {
							newPattern[mod + 0 + q * nWallSize] = result[0 + q * 3]
							newPattern[mod + 1 + q * nWallSize] = result[1 + q * 3]
							newPattern[mod + 2 + q * nWallSize] = result[2 + q * 3]
						}
					}
				}
				wallSize = nWallSize
			} else {
				regionCount := wallSize / 3
				nWallSize := regionCount * 4
				newPattern = make([]rune, nWallSize * nWallSize)

				for j := 0; j < regionCount; j++ {
					for k := 0; k < regionCount; k++ {
						mod := j * wallSize * 3 + k * 3
						stage := ThreeState{
							int(pattern[mod + 0] - '0'), int(pattern[mod + 1] - '0'), int(pattern[mod + 2] - '0'),
							int(pattern[mod + 0 + wallSize] - '0'), int(pattern[mod + 1 + wallSize] - '0'), int(pattern[mod + 2 + wallSize] - '0'),
							int(pattern[mod + 0 + wallSize * 2] - '0'), int(pattern[mod + 1 + wallSize * 2] - '0'), int(pattern[mod + 2 + wallSize * 2] - '0'),
						}
						result, _ := trans3[stage]

						mod = j * nWallSize * 4 + k * 4

						for q := 0; q < 4; q++ {
							newPattern[mod + 0 + q * nWallSize] = result[0 + q * 4]
							newPattern[mod + 1 + q * nWallSize] = result[1 + q * 4]
							newPattern[mod + 2 + q * nWallSize] = result[2 + q * 4]
							newPattern[mod + 3 + q * nWallSize] = result[3 + q * 4]
						}
					}
				}

				wallSize = nWallSize
			}

			pattern = newPattern
		}

		return string(pattern)
	}

	result5 := iterator(startPattern, 3, 5)
	result18 := iterator(startPattern, 3, 18)

	log.Printf("On: %v", strings.Count(result5, "1"))
	log.Printf("On: %v", strings.Count(result18, "1"))
	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-21")
	day21(input)
}
