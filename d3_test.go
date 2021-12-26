package aoc21

import (
	"testing"
)

func TestD3(t *testing.T) {
	casesP1 := []struct {
		in   string
		want int
	}{
		{`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, 198},
	}
	casesP2 := []struct {
		in   string
		want int
	}{
		{``, 0},
	}
	part1 := D3p1
	part2 := D3p2
	inputFile := "D3"

	Checker(t, casesP1, part1)
	Solution(inputFile, part1)

	Checker(t, casesP2, part2)
	Solution(inputFile, part2)
}
