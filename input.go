package main

import (
	"bufio"
	"github.com/imuli/newline/term"
	"os"
)

var line *Line = NewLine()
var reader = term.NewTermReader(bufio.NewReader(os.Stdin))
var end rune

type keyHandler func(r rune) keyHandler

func handleNormal(r rune) keyHandler {
	switch r {
	case term.Left:
		line.Shift(-1)
	case term.Right:
		line.Shift(1)
	case term.Home:
		line.Shift(line.ToEnd(0))
	case term.End:
		line.Shift(line.ToEnd(1))

	case term.Backspace:
		line.Delete(-1)
	case term.Delete:
		line.Delete(1)

	case term.Control:
		return handleControl
	case term.Meta:
		return handleMeta
	default:
		line.Insert([]rune{r})
	}
	return handleNormal
}

func handleControl(r rune) keyHandler {
	switch r {
	case 'b':
		line.Shift(-1)
	case 'f':
		line.Shift(1)
	case term.Left:
		line.Shift(line.Words(-1))
	case term.Right:
		line.Shift(line.Words(1))
	case 'a':
		line.Shift(line.ToEnd(0))
	case 'e':
		line.Shift(line.ToEnd(1))

		/* These all should be kills. */
	case 'h':
		line.Delete(-1)
	case 'w':
		line.Delete(line.Words(-1))
	case 'u':
		line.Delete(line.ToEnd(0))
	case 'k':
		line.Delete(line.ToEnd(1))

	case 'm':
		end = Newline
	case 'c':
		line.Clear()
		end = Interrupt
	case 'd':
		end = EndOfTransmission

	case 'v':
		nextRune, _, _ := reader.ReadRune()
		line.Insert([]rune{nextRune})
	}
	return handleNormal
}

func handleMeta(r rune) keyHandler {
	switch r {
	case 'b':
		line.Shift(line.Words(-1))
	case 'f':
		line.Shift(line.Words(1))
	}
	return handleNormal
}

var handle keyHandler

func CopyLine() (cont bool) {
	end = 0
	handle = handleNormal
	for loop := true; loop; {
		r, err := reader.ReadTermRune()
		if err != nil {
			return false
		}
		handle = handle(r)
		if end != 0 {
			cont = line.Flush(end)
			loop = false
		}
		line.Redraw()
	}
	OutputFlush()
	return
}
