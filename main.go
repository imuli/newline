package main

import (
	"flag"
	"syscall"
	"github.com/imuli/newline/term"
)

var termState term.State
var lines = 1.0

func main() {
	if termState, err := term.GetState(0); err == nil {
		tempState := *termState
		tempState.MakeRaw()
		tempState.Oflag |= syscall.OPOST
		term.SetState(0, &tempState)
		defer term.SetState(0, termState)
	}
	flag.Parse()

	for CopyLine() == nil {
		lines--
		if lines <= 0 {
			break
		}
	}
}
