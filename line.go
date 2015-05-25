package main

import ()

type Line struct {
	buffer []rune
	offset int
	cursor int
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
	if(mag < 0){
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
