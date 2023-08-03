package main

import (
	"fmt"

	"github.com/encleine/aoc-22/helper"
)

func solver(rd *helper.ReadLine, cargo *cargo) {
	val, _0 := rd.Next()
	for _0 {
		val, _0 = rd.Next()
		if len(val) == 0 {
			break
		}
		cargo.Step(val)
	}
	fmt.Println(cargo.Dump())
}
func main() {
	rd := helper.ReadInput("./input.txt")
	cargo := makeCargo(rd)
	solver(rd, &cargo)
}
