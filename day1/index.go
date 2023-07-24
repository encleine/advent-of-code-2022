package main

import (
	// "fmt"
	"math"
	"strconv"

	"github.com/encleine/aoc-22/helper"
)

var min, max, sum float64
var top3 [3]float64

func part1(rd *helper.ReadLine) {
	var val string
	_0 := true

	for _0 {
		val, _0 = rd.Next()
		cals, _ := strconv.ParseFloat(val, 64)
		if len(val) == 0 {

			max = math.Max(max, min)
			min = 0
			continue
		}
		min += cals
	}
}

func part2(rd *helper.ReadLine) {
	var val string
	_0 := true

	for _0 {
		val, _0 = rd.Next()
		cals, _ := strconv.ParseFloat(val, 64)

		if len(val) == 0 {

			for i, v := range top3 {

				if v < min {
					if i != 2 {
						copy(top3[i+1:3], top3[i:i+1])
					}
					top3[i] = min
					break
				}
			}

			min = 0
			continue
		}
		min += cals
	}

	for _, v := range top3 {
		sum += v
	}
}

func main() {
	rd := helper.ReadInput("./input.txt")
	defer rd.Close()
	// part1(rd)
	// part2(rd)
	// fmt.Println(max)
	// fmt.Println(sum)
}
