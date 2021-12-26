package aoc21

import (
	"io"
	"os"
	"strings"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	return b.String()
}
