package main

import (
	"bufio"
	"github.com/imuli/newline/term"
	"os"
)

var Line []rune = make([]rune, 0, 256)
var reader = term.NewTermReader(bufio.NewReader(os.Stdin))

func CopyLine() error {
	for loop := true; loop; {
		r, err := reader.ReadTermRune()
		if err != nil {
			return err
		}
		Line = append(Line, r)
		if r == term.Control {
			FlushLine()
			loop = false
		}
		RedrawLine()
		OutputFlush()
	}
	return nil
}
