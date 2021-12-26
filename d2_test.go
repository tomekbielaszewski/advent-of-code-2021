package aoc21

import (
	"testing"
)

func TestD2(t *testing.T) {
	casesP1 := []struct {
		in   string
		want int
	}{
		{`forward 5
down 5
forward 8
up 3
down 8
forward 2`, 150},
	}
	casesP2 := []struct {
		in   string
		want int
	}{
		{`forward 5
down 5
forward 8
up 3
down 8
forward 2`, 900},
	}
	part1 := D2p1
	part2 := D2p2
	inputFile := "D2"

	Checker(t, casesP1, part1)
	Solution(inputFile, part1)

	Checker(t, casesP2, part2)
	Solution(inputFile, part2)
}
