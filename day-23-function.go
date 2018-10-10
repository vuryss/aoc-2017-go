package main

import "log"

func cycle() int {
	b, c, d, e, f, g, h := 108400, 125400, 2, 2, 1, 0, 0

	for {
		g = d * e - b

		if g == 0 {
			f = 0
		}

		if f != 0 &&  b % d == 0 {
			e = b / d
		} else {
			e = b
		}

		g = e - b

		if g == 0 {
			d += 1
			g = d - b

			if g == 0 {
				if f == 0 {
					h += 1
				}

				g = b - c

				if g == 0 {
					return h
				}

				b += 17
				f = 1
				d = 2
			}

			e = 2
		}
	}
}

func main() {
	log.Printf("H value: %v", cycle())
}