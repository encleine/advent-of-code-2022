package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/encleine/aoc-22/helper"
)

var count int

func solver(rd *helper.ReadLine, sections func(x, y *[]int) int) {
	var val string
	_0 := true
	for _0 {
		val, _0 = rd.Next()
		if len(val) == 0 {
			break
		}
		ranges := strings.Split(val, ",")

		x := strings.Split(ranges[0], "-")
		y := strings.Split(ranges[1], "-")

		count += sections(_IntSlice(&x), _IntSlice(&y))
	}
}

func part1(x, y *[]int) int {
	switch {
	case (*y)[0] < (*x)[0] && (*y)[1] < (*x)[1]:
		fallthrough
	case (*x)[0] < (*y)[0] && (*x)[1] < (*y)[1]:
		return 0

	}
	return 1
}

func part2(x, y *[]int) int {
	switch {
	case (*y)[1] < (*x)[0] ||
		(*y)[0] > (*x)[1]:
		return 0
	}
	return 1
}
func _IntSlice(ar *[]string) *[]int {
	var sl []int
	for _, v := range *ar {
		val, _ := strconv.Atoi(v)
		sl = append(sl, val)
	}
	return &sl
}

func main() {
	rd := helper.ReadInput("./input.txt")
	defer rd.Close()
	// solver(rd, part1 || part2)
	// fmt.Println("sum of the priorities ", count)
}
