package problem2215

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {4, 6}},
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			want: [][]int{{3}, {}},
		},
		{
			name: "Example 3",
			args: args{
				nums1: []int{2, 5},
				nums2: []int{2, 3},
			},
			want: [][]int{{5}, {3}},
		},
		{
			name: "Example 4",
			args: args{
				nums1: []int{1},
				nums2: []int{4, 5},
			},
			want: [][]int{{1}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)

			got1 := got[0]
			got2 := got[1]

			sort.Ints(got1)
			sort.Ints(got2)

			got = [][]int{got1, got2}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
