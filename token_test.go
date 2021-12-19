package lad

import (
	"testing"
	"unicode"

	"github.com/mozillazg/go-pinyin"
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

func TestPinyin(t *testing.T) {
	hans := "重"
	t.Log(unicode.Is(unicode.Han, []rune(hans)[0]))
	// 默认
	a := pinyin.NewArgs()
	a.Heteronym = true
	t.Log(pinyin.Pinyin(hans, a))
}
