package aoc21

import (
	"fmt"
	"testing"
)

func TestD1p1(t *testing.T) {
	cases := []struct {
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
	for _, c := range cases {
		got := D1p1(c.in)
		if got != c.want {
			t.Errorf("solution of (%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestD1p1solution(t *testing.T) {
	file := "D1"
	in := Read(file)
	fmt.Printf("%s solution: %d\n", file, D1p1(in))
}

func TestD1p2(t *testing.T) {
	cases := []struct {
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
	for _, c := range cases {
		got := D1p2(c.in)
		if got != c.want {
			t.Errorf("solution of (%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestD1p2solution(t *testing.T) {
	file := "D1"
	in := Read(file)
	fmt.Printf("%s solution: %d\n", file, D1p2(in))
}
