package aoc21

import (
	"strconv"
	"strings"
)

func D2p1(input string) int {
	els := strings.Split(input, "\n")
	horizontal := 0
	depth := 0
	for _, el := range els {
		el = strings.TrimSpace(el)
		lineParts := strings.Split(el, " ")
		command := lineParts[0]
		val := lineParts[1]
		value, _ := strconv.Atoi(val)

		switch {
		case command == "forward":
			horizontal += value
		case command == "down":
			depth += value
		case command == "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func D2p2(input string) int {
	els := strings.Split(input, "\n")
	aim := 0
	horizontal := 0
	depth := 0
	for _, el := range els {
		el = strings.TrimSpace(el)
		lineParts := strings.Split(el, " ")
		command := lineParts[0]
		val := lineParts[1]
		value, _ := strconv.Atoi(val)

		switch {
		case command == "forward":
			horizontal += value
			depth += aim * value
		case command == "down":
			aim += value
		case command == "up":
			aim -= value
		}
	}
	return horizontal * depth
}
