package term

import (
	"bufio"
	"unicode/utf8"
)

const (
	Up      = 'U'
	Down    = 'D'
	Right   = 'R'
	Left    = 'L'
	Meta    = 'M'
	Control = 'C'
	Shift   = 'S'
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

func (t *TermReader) pushRunes() error {
	r, _, err := t.ReadRune()
	if err != nil {
		return err
	}
	if r == '\x1b' {
		return t.readEscape()
	}
	if r < ' ' {
		t.q = append(t.q, Control)
		r |= '`'
	}
	t.q = append(t.q, r)
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
		//		return readESCO(b)
	case 'o':
		//		return readESCo(b)
	}
	t.UnreadRune()
	t.q = append(t.q, Meta)
	return nil
}

func (t *TermReader) readPS() (n int, r rune, err error) {
	n = 0
	for r = '0'; '0' <= r && r <= '9' && err == nil; r, _, err = t.ReadRune() {
		n = n*10 + int(r-'0')
	}
	return
}

func (t *TermReader) readCSI() (err error) {
	_, r, err := t.readPS()
	switch r {
	case 'A':
		t.q = append(t.q, Up)
	case 'B':
		t.q = append(t.q, Down)
	case 'C':
		t.q = append(t.q, Right)
	case 'D':
		t.q = append(t.q, Left)
	case 'a':
		t.q = append(t.q, Shift, Up)
	case 'b':
		t.q = append(t.q, Shift, Down)
	case 'c':
		t.q = append(t.q, Shift, Right)
	case 'd':
		t.q = append(t.q, Shift, Left)
	case 'Z':
		t.q = append(t.q, Shift, Control, 'i')
	}
	return
}
