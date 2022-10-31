package problem_13

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with 'III' input",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "Test case with 'LVIII' input",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "Test case with 'MCMXCIV' input",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
