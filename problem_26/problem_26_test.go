package problem26

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 3, 6, 8, 8, 9, 9, 9, 12, 18},
			},
			want: 6,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{
				nums: []int{1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
