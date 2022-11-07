package problem_1323

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with num = 9669 input",
			args: args{
				num: 9669,
			},
			want: 9969,
		},
		{
			name: "Test case with num = 9996 input",
			args: args{
				num: 9996,
			},
			want: 9999,
		},
		{
			name: "Test case with num = 9999 input",
			args: args{
				num: 9999,
			},
			want: 9999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
