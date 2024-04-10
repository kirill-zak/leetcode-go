package problem349

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
		{
			name: "Example 3",
			args: args{
				nums1: []int{1, 4},
				nums2: []int{},
			},
			want: []int{},
		},
		{
			name: "Example 4",
			args: args{
				nums1: []int{},
				nums2: []int{4, 6},
			},
			want: []int{},
		},
		{
			name: "Example 5",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.args.nums1, tt.args.nums2)

			sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
