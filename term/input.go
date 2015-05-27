package term

import (
	"bufio"
	"unicode/utf8"
)

const (
	Backspace = 0x7f
	Normal = 0xd800 + iota // surrogates are not valid UTF-8 so are safe for internal use
	Control
	Meta
	Shift
	Up
	Down
	Left
	Right
	Prior
	Next
	Home
	End
	Find
	Select
	Insert
	Delete
	Function
)

type TermReader struct {
	r bufio.Reader
	q []rune
}

func NewTermReader(r *bufio.Reader) (t *TermReader) {
	t = new(TermReader)
	t.r = *r
	t.q = make([]rune, 0, 4)
	return t
}

// ReadTermRune reads escape sequence and runes as keystrokes
func (t *TermReader) ReadTermRune() (r rune, err error) {
	if len(t.q) == 0 {
		err = t.pushRunes()
		if len(t.q) == 0 {
			return utf8.RuneError, err
		}
	}
	r = t.q[0]
	t.q = t.q[1:len(t.q)]
	return r, nil
}

func (t *TermReader) ReadRune() (rune, int, error) {
	return t.r.ReadRune()
}

func (t *TermReader) UnreadRune() error {
	return t.r.UnreadRune()
}

func (t *TermReader) push(it ...rune) {
	t.q = append(t.q, it...)
}

func (t *TermReader) pushRunes() error {
	r, _, err := t.ReadRune()
	if err != nil {
		return err
	}
	if r == '\x1b' {
		return t.readEscape()
	}
	if r < ' ' {
		t.push(Control)
		r |= '`'
	}
	t.push(r)
	return nil
}

func (t *TermReader) readEscape() error {
	r, _, err := t.ReadRune()
	if err != nil {
		return err
	}
	switch r {
	case '[':
		return t.readCSI()
	case 'O':
		return t.readEscO()
	}
	t.UnreadRune()
	t.push(Meta)
	return nil
}

func (t *TermReader) readPS() (n int, r rune, err error) {
	n = 0
	for r = '0'; '0' <= r && r <= '9' && err == nil; r, _, err = t.ReadRune() {
		n = n*10 + int(r-'0')
	}
	return
}

func (t *TermReader) readCSI() error {
	n, r, err := t.readPS()
	switch r {
	case 'A':
		t.push(Up)
	case 'B':
		t.push(Down)
	case 'C':
		t.push(Right)
	case 'D':
		t.push(Left)
	case 'a':
		t.push(Shift, Up)
	case 'b':
		t.push(Shift, Down)
	case 'c':
		t.push(Shift, Right)
	case 'd':
		t.push(Shift, Left)
	case 'Z':
		t.push(Shift, Control, 'i')
	case '@':
		t.push(Control, Shift, fKey(n))
	case '^':
		t.push(Control, fKey(n))
	case '$':
		t.push(Shift, fKey(n))
	case '~':
		t.push(fKey(n))
	}
	return err
}

func (t *TermReader) readEscO() error {
	r, _, err := t.ReadRune()
	switch r {
	case 'a':
		t.push(Control, Up)
	case 'b':
		t.push(Control, Down)
	case 'c':
		t.push(Control, Right)
	case 'd':
		t.push(Control, Left)
	case 'A':
		t.push(Control, Shift, Up)
	case 'B':
		t.push(Control, Shift, Down)
	case 'C':
		t.push(Control, Shift, Right)
	case 'D':
		t.push(Control, Shift, Left)
	}
	return err
}

var fkeys = [...]rune{
	Normal, Find, Insert, Delete, Select, Prior, Next, Home, End, Normal, Normal,
	Normal, Function + 1, Function + 2, Function + 3, Function + 4, Function + 5,
	Normal, Function + 6, Function + 7, Function + 8, Function + 9, Function + 10,
	Normal, Function + 11, Function + 12, Function + 13, Function + 14,
	Normal, Function + 15, Function + 16,
	Normal, Function + 17, Function + 18, Function + 19, Function + 20,
}

func fKey(n int) rune {
	if n < len(fkeys) {
		return fkeys[n]
	}
	return Normal
}
