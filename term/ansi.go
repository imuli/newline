package term

import (
	"fmt"
	runewidth "github.com/mattn/go-runewidth"
	"io"
)

func CursorShift(writer io.Writer, n int) {
	if n == 0 {
		return
	}
	var c = 'C'
	if n < 0 {
		c = 'D'
		n = -n
	}
	fmt.Fprintf(writer, "\x1b[%d%c", n, c)
}

func EraseLine(writer io.Writer) {
	fmt.Fprint(writer, "\x1b[K")
}

func DisplayWidth(text []rune) (sum int){
	for _, r := range text {
		sum += runewidth.RuneWidth(r)
	}
	return
}
