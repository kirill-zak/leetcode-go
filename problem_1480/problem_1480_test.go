package problem_1480

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

type testCase struct {
	name string
	args args
	want []int
}

func Test_runningSum(t *testing.T) {
	tests := makeTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runningSumWithSum(t *testing.T) {
	tests := makeTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSumWithSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSumWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeTestCases() []testCase {
	return []testCase{
		{
			name: "Test case for [1,2,3,4] input",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "Test case for [1,1,1,1,1] input",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case for [3,1,2,10,1] input",
			args: args{
				nums: []int{3, 1, 2, 10, 1},
			},
			want: []int{3, 4, 6, 16, 17},
		},
	}
}
