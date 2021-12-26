package aoc21

import (
	"strconv"
	"strings"
)

func D1p1(input string) int {
	els := strings.Split(input, "\n")
	ddepth := 0
	prev := 0
	for i, el := range els {
		el = strings.TrimSpace(el)
		elint, _ := strconv.Atoi(el)
		if i > 0 {
			if prev < elint {
				ddepth++
			}
		}
		prev = elint
	}
	return ddepth
}

func D1p2(input string) int {
	els := strings.Split(input, "\n")
	ddepth := 0
	for i, _ := range els {
		if i > 2 {
			window1 := sum(els[i-3 : i])
			window2 := sum(els[i-2 : i+1])
			if window1 < window2 {
				ddepth++
			}
		}
	}
	return ddepth
}

func sum(els []string) int {
	sum := 0
	for _, el := range els {
		el = strings.TrimSpace(el)
		elint, _ := strconv.Atoi(el)
		sum += elint
	}
	return sum
}
