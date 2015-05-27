package main

import (
	"unicode"
)

type Line struct {
	buffer []rune
	offset int
}

func NewLine() *Line {
	line := new(Line)
	line.buffer = make([]rune, 0, 256)
	return line
}

// Insert runes into the line.
func (l *Line) Insert(text []rune) {
	l.buffer = l.buffer[:len(l.buffer)+len(text)]
	copy(l.buffer[l.offset+len(text):], l.buffer[l.offset:])
	copy(l.buffer[l.offset:l.offset+len(text)], text[:])
	l.offset += len(text)
}

func (l *Line) Delete(mag int) {
	if mag < 0 {
		l.offset += mag
		mag = -mag
	}
	copy(l.buffer[l.offset:], l.buffer[l.offset+mag:])
	l.buffer = l.buffer[:len(l.buffer)-mag]
}

func (l *Line) Clear() {
	l.buffer = l.buffer[0:0]
	l.offset = 0
}

func (l *Line) Shift(n int) {
	l.offset += n
	switch {
	case l.offset < 0:
		l.offset = 0
	case l.offset > len(l.buffer):
		l.offset = len(l.buffer)
	}
}

func (l *Line) ToEnd(n int) int {
	return n*len(l.buffer) - l.offset
}

func (l *Line) Words(n int) int {
	var direction int
	i := l.offset
	switch {
	case n < 0:
		direction = -1
		i--
		n = -n
	case n > 0:
		direction = 1
	default:
		return 0
	}
	for ; n > 0; n-- {
		for i >= 0 && i < len(line.buffer) && unicode.IsSpace(line.buffer[i]) {
			i += direction
		}
		for i >= 0 && i < len(line.buffer) && !unicode.IsSpace(line.buffer[i]) {
			i += direction
		}
	}
	if direction == -1 {
		i++
	}
	return i - l.offset
}
