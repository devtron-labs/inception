package pkg

import "testing"

func TestStripQuotes(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `test for "`,
			args: args{`"hello"`},
			want: "hello",
		},
		{
			name: `test for ''`,
			args: args{`'hello'`},
			want: "hello",
		},
		{
			name: "test for `",
			args: args{"`hello`"},
			want: "hello",
		},
		{
			name: `test for without quote`,
			args: args{`hello`},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripQuotes(tt.args.text); got != tt.want {
				t.Errorf("StripQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
