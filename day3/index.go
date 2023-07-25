package main

import (
	"fmt"

	"github.com/encleine/aoc-22/helper"
)

var sum int

func normalize(ru rune) int {
	var holder rune
	switch {
	case ru >= 'a' && 'z' >= ru:
		holder = ru - 96
	case ru >= 'A' && 'Z' >= ru:
		holder = ru - 38
	}
	return int(holder)
}

func main() {
	rd := helper.ReadInput("./input.txt")
	defer rd.Close()
	part2Solver(rd)
	fmt.Println("sum of the priorities ", sum)

}
