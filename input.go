package main

import (
	"bufio"
	"github.com/imuli/newline/term"
	"os"
)

var line *Line = NewLine()
var reader = term.NewTermReader(bufio.NewReader(os.Stdin))

func CopyLine() error {
	for loop := true; loop; {
		r, err := reader.ReadTermRune()
		if err != nil {
			return err
		}
		line.Insert([]rune{r})
		if r == term.Control {
			line.Flush()
			loop = false
		}
		line.Redraw()
		OutputFlush()
	}
	return nil
}
