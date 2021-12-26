package aoc21

import (
	"testing"
)

func TestD1(t *testing.T) {
	casesP1 := []struct {
		in   string
		want int
	}{
		{`199
200
208
210
200
207
240
269
260
263`, 7},
	}
	casesP2 := []struct {
		in   string
		want int
	}{
		{`199
200
208
210
200
207
240
269
260
263`, 5},
	}
	part1 := D1p1
	part2 := D1p2
	inputFile := "D1"

	Checker(t, casesP1, part1)
	Solution(inputFile, part1)

	Checker(t, casesP2, part2)
	Solution(inputFile, part2)
}
