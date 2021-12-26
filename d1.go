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
