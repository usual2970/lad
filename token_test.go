package lad

import (
	"testing"
)

func Test_token_next(t1 *testing.T) {
	raw := "ab    阿宾   cdadfadfadfad"
	type fields struct {
		input []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TOKEN_NEXT",
			fields: fields{
				input: []rune(raw),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := newToken(string(tt.fields.input))

			for rs := t.next(); rs != ""; rs = t.next() {
				t1.Log(rs, t.index)
			}
		})
	}
}

func Test_token_nextPinyin(t1 *testing.T) {
	raw := "ab    啊中重   cdadfadfadfad"
	type fields struct {
		input []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TOKEN_NEXT",
			fields: fields{
				input: []rune(raw),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := newToken(string(tt.fields.input))

			for rs := t.nextPinyin(); rs != nil; rs = t.nextPinyin() {
				t1.Log(rs, t.index)
			}
		})
	}
}

func Test_token_buildGraph(t1 *testing.T) {

	raw := "ab    啊中重   cdadfadfadfad"

	t := newToken(raw)

	graph := t.buildGraph()

	graph.dfs()
}
