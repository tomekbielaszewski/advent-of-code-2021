package aoc21

import (
	"fmt"
	"testing"
)

type Solver func(string) int

func TestExample(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{``, 0},
	}
	part1 := Dxp1
	part2 := Dxp2
	inputFile := "Dx"

	Checker(t, cases, part1)
	Solution(inputFile, part1)

	Checker(t, cases, part2)
	Solution(inputFile, part2)
}

func Checker(t *testing.T, cases []struct {
	in   string
	want int
}, function Solver) {
	for _, c := range cases {
		got := function(c.in)
		if got != c.want {
			t.Errorf("solution of (%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Solution(inputFile string, function Solver) {
	in := Read(inputFile)
	solution := function(in)
	fmt.Printf("%s solution: %d\n", inputFile, solution)
}
