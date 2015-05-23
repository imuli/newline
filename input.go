package main

import (
	"bufio"
	"os"
)

var Line []rune = make([]rune, 0, 256)
var reader = bufio.NewReader(os.Stdin)

func CopyLine() error {
	for loop := true;loop; {
		r, _, err := reader.ReadRune()
		if err != nil {
			return err
		}
		Line = append(Line, r)
		if r == '\n' {
			FlushLine()
			loop = false
		}
		RedrawLine()
		OutputFlush()
	}
	return nil
}
