package main

import (
	"strings"

	"github.com/encleine/aoc-22/helper"
)

func part2Solver(rd *helper.ReadLine) {
	var line1, line2, line3 string
	_0 := true

	for _0 {
		line1, _ = rd.Next()
		line2, _ = rd.Next()
		line3, _0 = rd.Next()
		if len(line1) == 0 {
			break
		}
		sum += groups(line1, line2, line3)
	}
}
func groups(line1, line2, line3 string) int {

	for _, v := range line1 {
		if strings.ContainsRune(line2, v) {
			if strings.ContainsRune(line3, v) {
				return normalize(v)
			}
		}
	}
	return 0
}
