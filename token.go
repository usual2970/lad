package lad

import "unicode"

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
	t.moveNext()
	return string(current)
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
