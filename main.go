package main

import (
	"flag"
	"github.com/imuli/newline/term"
	"strconv"
	"syscall"
)

var termState term.State

func main() {
	var lines float64 = 1.0
	if termState, err := term.GetState(0); err == nil {
		tempState := *termState
		tempState.MakeRaw()
		tempState.Oflag |= syscall.OPOST
		term.SetState(0, &tempState)
		defer term.SetState(0, termState)
	}
	flag.Parse()
	if p := flag.Args(); len(p) > 0 {
		var err error
		lines, err = strconv.ParseFloat(p[len(p)-1], 64)
		if err != nil {
			return
		}
	}

	for lines > 0 && CopyLine() {
		lines--
	}
}
