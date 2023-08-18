package main

import (
	"fmt"

	"github.com/encleine/aoc-22/helper"
)

type buffer struct {
	buffer string
	start  int
	end    int
	space  int
}

func (b *buffer) String() string {
	return b.buffer[b.start:b.end]
}

func (b *buffer) HasDup() bool {
	temp := map[rune]bool{}
	for _, v := range b.String() {
		if _, ok := temp[v]; ok {
			return true
		}
		temp[v] = false
	}
	return false
}
func (b *buffer) FistMarker() int {
	for i := 0; i < len(b.buffer)-b.space; i++ {
		if !b.HasDup() {
			return b.end
		}
		b.start++
		b.end++
	}
	return b.end
}
func build(s string, space int) *buffer {
	return &buffer{
		buffer: s,
		start:  0,
		end:    space,
		space:  space,
	}
}
func solver(rd *helper.ReadLine) {
	val, _0 := rd.Next()
	for _0 {
		if len(val) == 0 {
			break
		}
		fmt.Printf("%d-", build(val, 14).FistMarker())

		val, _0 = rd.Next()
	}
}
func main() {
	rd := helper.ReadInput("./input.txt")

	solver(rd)
}
