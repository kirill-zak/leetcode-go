package problem_1672

import "testing"

type args struct {
	accounts [][]int
}

type testCase struct {
	name string
	args args
	want int
}

func Test_maximumWealth(t *testing.T) {
	tests := makeTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumWealthWithChannel(t *testing.T) {
	tests := makeTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWealth := maximumWealthWithChannel(tt.args.accounts); gotWealth != tt.want {
				t.Errorf("maximumWealthWithChannel() = %v, want %v", gotWealth, tt.want)
			}
		})
	}
}

func makeTestCases() []testCase {
	return []testCase{
		{
			name: "Test case with [[1,2,3],[3,2,1]] input",
			args: args{
				accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			},
			want: 6,
		},
		{
			name: "Test case with [[1,5],[7,3],[3,5]] input",
			args: args{
				accounts: [][]int{{1, 5}, {7, 3}, {3., 5}},
			},
			want: 10,
		},
		{
			name: "Test case with [[2,8,7],[7,1,3],[1,9,5]] input",
			args: args{
				accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			},
			want: 17,
		},
	}
}
