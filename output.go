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

const (
	Newline           = 0x0a
	Interrupt         = 0x03
	EndOfTransmission = 0x04
)

func Render(writer *bufio.Writer, text []rune) {
	for i := range text {
		writer.WriteRune(text[i])
	}
}

var cursor int

func (l *Line) Redraw() {
	term.CursorShift(terminal, -cursor)
	Render(terminal, l.buffer)
	term.EraseLine(terminal)
	term.CursorShift(terminal, -term.DisplayWidth(l.buffer[l.offset:]))
	cursor = term.DisplayWidth(l.buffer[:l.offset])
	terminal.Flush()
}

func (l *Line) Flush(r rune) bool {
	Render(output, l.buffer)
	cont := true
	switch {
	case *ptyMode, r == Newline:
		output.WriteRune(r)
	case r == Interrupt, r == EndOfTransmission && len(l.buffer) == 0:
		cont = false
	}
	l.Clear()
	return cont
}

func OutputFlush() {
	output.Flush()
}
