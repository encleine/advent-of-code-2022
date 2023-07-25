package main

import (
	"strings"

	"github.com/encleine/aoc-22/helper"
)

func Par1Solver(rd *helper.ReadLine) {
	var val string
	_0 := true

	for _0 {
		val, _0 = rd.Next()
		if len(val) == 0 {
			break
		}
		sum += rucksack(val)
	}
}
func rucksack(line string) int {
	for _, v := range line[:len(line)/2] {
		if strings.ContainsRune(line[len(line)/2:], v) {
			return normalize(v)
		}
	}
	return 0
}
