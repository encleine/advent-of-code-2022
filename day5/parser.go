package main

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"

	"github.com/encleine/aoc-22/helper"
)

// using string for because i'll convert them to strings later anyways
type stack []string

func (st *stack) Push(char string) {
	*st = append(*st, char)
}

func (st *stack) Pop(other *stack) {
	other.Push((*st)[len(*st)-1])
	*st = (*st)[:len(*st)-1]
}
func (st *stack) Peek() string {
	return (*st)[len(*st)-1]
}

var hasNum = regexp.MustCompile("[0-1]")

type cargo struct {
	Crates []string
	Stacks []stack
}

func makeCargo(rd *helper.ReadLine) cargo {
	var car cargo

	for {
		val, _ := rd.Next()
		if hasNum.Match([]byte(val)) {
			car.clean(val)
			break
		}
		car.Crates = append(car.Crates, val)
	}
	car.parseCrates()
	return car
}

func (cr *cargo) clean(val string) {
	var leng int
	for _, v := range val {
		switch v {
		case ' ':
			continue
		default:
			leng++
		}
	}
	cr.Stacks = make([]stack, leng)
}

func (cr *cargo) parseCrates() {
	stacks := cr.Stacks
	crates := cr.Crates
	count := len(crates) - 1

	for i := count; i >= 0; i-- {
		carte := crates[i]
		for in := 0; in < len(stacks); in++ {
			if c := carte[in*4+1]; c != ' ' {
				stacks[in].Push(string(c))
			}
		}
	}
	cr.Crates = nil
}

func (cr *cargo) Dump() string {
	var buffer bytes.Buffer
	stacks := cr.Stacks
	for _, car := range stacks {
		buffer.WriteString(car.Peek())
	}
	return buffer.String()
}

func parseInt(line *string, start int) int {
	var buffer bytes.Buffer
	for len(*line) != start && unicode.IsNumber(rune((*line)[start])) {
		buffer.WriteByte((*line)[start])
		start++
	}
	num, _ := strconv.Atoi(buffer.String())
	return num
}
func (cr *cargo) Mover9000(amount, from, to int) {
	stacks := cr.Stacks
	for i := 0; i < amount; i++ {
		stacks[from-1].Pop(&stacks[to-1])
	}
}
func (cr *cargo) Mover9001(amount, from, to int) {
	temp := new(stack)
	stacks := cr.Stacks
	// lazy way to solve this
	for i := 0; i < amount; i++ {
		stacks[from-1].Pop(temp)
	}
	for i := 0; i < amount; i++ {
		temp.Pop(&stacks[to-1])
	}
}
func (cr *cargo) Step(line string) {
	var amount, from, to int
	for i := 0; i < len(line); i++ {
		switch let := line[i]; let {
		case 'm':
			i += 5
			amount = parseInt(&line, i)
		case 'f':
			i += 5
			from = parseInt(&line, i)
		case 't':
			i += 3
			to = parseInt(&line, i)
		}
	}
	cr.Mover9001(amount, from, to)
}
