package aoc21

import (
	"fmt"
	"testing"
)

func TestD2p1(t *testing.T) {
	cases := []struct {
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
	for _, c := range cases {
		got := D2p1(c.in)
		if got != c.want {
			t.Errorf("solution of (%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestD2p1solution(t *testing.T) {
	file := "D2"
	in := Read(file)
	solution := D2p1(in)
	fmt.Printf("%s solution: %d\n", file, solution)
}

func TestD2p2(t *testing.T) {
	cases := []struct {
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
	for _, c := range cases {
		got := D2p2(c.in)
		if got != c.want {
			t.Errorf("solution of (%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestD2p2solution(t *testing.T) {
	file := "D2"
	in := Read(file)
	solution := D2p2(in)
	fmt.Printf("%s solution: %d\n", file, solution)
}
