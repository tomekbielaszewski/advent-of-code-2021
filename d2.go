package aoc21

import (
	"strconv"
	"strings"
)

func D2p1(input string) int {
	els := strings.Split(input, "\n")
	x := 0
	y := 0
	for _, el := range els {
		el = strings.TrimSpace(el)
		lineParts := strings.Split(el, " ")
		command := lineParts[0]
		val := lineParts[1]
		value, _ := strconv.Atoi(val)

		switch {
		case command == "forward":
			x += value
		case command == "down":
			y += value
		case command == "up":
			y -= value
		}
	}
	return x * y
}

func D2p2(input string) int {
	return 0
}
