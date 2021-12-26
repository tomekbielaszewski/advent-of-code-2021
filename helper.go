package aoc21

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		errorMsg := "error reading file!!"
		fmt.Println(errorMsg)
		fmt.Println(err)
		return errorMsg
	}
	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	return b.String()
}
