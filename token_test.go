package lad

import "testing"

func Test_token_next(t1 *testing.T) {
	raw := "hello world 我们都有一个家"
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
				t1.Log(rs,t.index)
			}
		})
	}
}
