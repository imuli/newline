package main

import (
	"bufio"
	"flag"
	"github.com/imuli/newline/term"
	"os"
)

var ptyMode = flag.Bool("pty", false, "enable pty mode")
var cursor int
var terminal = bufio.NewWriter(os.Stderr)
var output = bufio.NewWriter(os.Stdout)

func Render(writer *bufio.Writer) {
	for i := range Line {
		writer.WriteRune(Line[i])
	}
}

func RedrawLine() {
	term.CursorShift(terminal, -cursor)
	Render(terminal)
	cursor = len(Line)
	term.EraseLine(terminal)
	terminal.Flush()
}

func FlushLine() {
	Render(output)
	Line = Line[0:0]
}

func OutputFlush() {
	output.Flush()
}
