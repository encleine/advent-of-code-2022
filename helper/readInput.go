package helper

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	// "strconv"
)

type ReadLine struct {
	scanner *bufio.Scanner
	file    *os.File
}

func (nl *ReadLine) Next() (string, bool) {
	sc := nl.scanner
	scan := sc.Scan()
	return sc.Text(), scan
}

func (nl *ReadLine) Close() {
	nl.file.Close()
}

func ReadInput(path string) *ReadLine {
	file, _0 := os.Open(path)
	if _0 != nil {
		log.Fatal(_0)
	}
	scanner := bufio.NewScanner(file)

	R := &ReadLine{
		scanner,
		file,
	}
	return R
}
