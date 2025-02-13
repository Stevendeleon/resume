package pkg

import "testing"

func Test_lenMinusOne(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty slice",
			args: args{
				s: []string{},
			},
			want: -1,
		},
		{
			name: "slice with one element",
			args: args{
				s: []string{"hello"},
			},
			want: 0,
		},
		{
			name: "slice with multiple elements",
			args: args{
				s: []string{"hello", "world", "test"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenMinusOne(tt.args.s); got != tt.want {
				t.Errorf("lenMinusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
