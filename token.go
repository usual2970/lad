package lad

import (
	"unicode"
)

type token struct {
	index  int
	origin []rune
	input  []rune
}

func newToken(raw string) *token {
	return &token{
		input:  []rune(raw),
		origin: []rune(raw),
		index:  -1,
	}
}

func (t *token) next() string {
	for unicode.IsSpace(t.current()) {
		t.moveNext()
	}

	current := t.current()
	if current == -1 {
		return ""
	}

	if !t.currentIsHan() {
		return t.readWord()
	}
	t.moveNext()
	return string(current)
}

func (t *token) readWord() string {
	word := []rune{}

	for {
		if unicode.IsSpace(t.current()) ||
			t.currentIsHan() || t.currentIs(-1) {
			return string(word)
		}

		word = append(word, t.current())
		t.moveNext()
	}
}

func (t *token) currentIsHan() bool {
	return unicode.Is(unicode.Han, t.current())
}

func (t *token) currentIs(runes ...rune) bool {
	for _, r := range runes {
		if r == t.current() {
			return true
		}
	}
	return false
}

func (t *token) moveNext() {
	t.index++
	t.input = t.input[1:]
}

func (t *token) prevNStr(index, preN int) string {
	rs := make([]rune, 0)
	for preN > 0 {
		r := t.origin[index]
		index--
		rs = append([]rune{r}, rs...)
		if unicode.IsSpace(r) {
			continue
		}
		preN--
	}
	return string(rs)
}

func (t *token) current() rune {
	return t.getRuneAt(0)
}

func (t *token) getRuneAt(i int) rune {
	if len(t.input) == i {
		return -1
	}

	return t.input[i]
}
