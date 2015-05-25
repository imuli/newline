package main

import (
	"bufio"
	"flag"
	"github.com/imuli/newline/term"
	"os"
)

var ptyMode = flag.Bool("pty", false, "enable pty mode")
var terminal = bufio.NewWriter(os.Stderr)
var output = bufio.NewWriter(os.Stdout)

func Render(writer *bufio.Writer, text []rune) {
	for i := range text {
		writer.WriteRune(text[i])
	}
}

func (l *Line) Redraw() {
	term.CursorShift(terminal, -l.cursor)
	Render(terminal, l.buffer)
	term.EraseLine(terminal)
	term.CursorShift(terminal, -term.DisplayWidth(l.buffer[l.offset:]))
	l.cursor = term.DisplayWidth(l.buffer[:l.offset])
	terminal.Flush()
}

func (l *Line) Flush() {
	Render(output, l.buffer)
	l.Clear()
}

func OutputFlush() {
	output.Flush()
}
