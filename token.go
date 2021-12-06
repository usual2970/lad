package lad

import "unicode"

type token struct {
	raw   string
	input []rune
}

func newToken(raw string) *token {
	return &token{
		raw:   raw,
		input: []rune(raw),
	}
}

func (t *token) next() string {
	for unicode.IsSpace(t.current()) {
		t.input = t.input[1:]
	}

	current := t.current()
	if current == -1 {
		return ""
	}

	t.input = t.input[1:]
	return string(current)
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
