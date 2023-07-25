package main

import (
	"fmt"
	"strings"

	"github.com/encleine/aoc-22/helper"
)

var score int

// opponent, player
type Rule map[[2]string]int

var rule1 = Rule{
	{"B", "X"}: 1,
	{"C", "Y"}: 2,
	{"A", "Z"}: 3,

	{"A", "X"}: 4,
	{"B", "Y"}: 5,
	{"C", "Z"}: 6,

	{"C", "X"}: 7,
	{"A", "Y"}: 8,
	{"B", "Z"}: 9,
}

var rule2 = Rule{
	{"B", "X"}: 1,
	{"C", "X"}: 2,
	{"A", "X"}: 3,

	{"A", "Y"}: 4,
	{"B", "Y"}: 5,
	{"C", "Y"}: 6,

	{"C", "Z"}: 7,
	{"A", "Z"}: 8,
	{"B", "Z"}: 9,
}

func solver(rd *helper.ReadLine, rule Rule) {
	var val string
	_0 := true

	for _0 {
		val, _0 = rd.Next()
		if len(val) == 0 {
			break
		}
		game := strings.Split(val, " ")

		score += rule[[2]string(game)]
	}
}
func main() {
	rd := helper.ReadInput("./input.txt")
	defer rd.Close()

	solver(rd, rule2)
	fmt.Println("final score is ", score)
}
