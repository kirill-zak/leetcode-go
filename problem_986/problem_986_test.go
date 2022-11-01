package problem_986

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		firstList  [][]int
		secondList [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test case with firstList = [[0,9],[20,26],[40,49]], secondList = [[0,1],[3,5],[7,8],[12," +
				"16],[17,18],[22,24],[25,26],[34,37],[39,42]]",
			args: args{
				firstList:  [][]int{{0, 9}, {20, 26}, {40, 49}},
				secondList: [][]int{{0, 1}, {3, 5}, {7, 8}, {12, 16}, {17, 18}, {22, 24}, {25, 26}, {34, 37}, {37, 42}},
			},
			want: [][]int{{0, 1}, {3, 5}, {7, 8}, {22, 24}, {25, 26}, {40, 42}},
		},
		{
			name: "Test case with firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]] input",
			args: args{
				firstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
				secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			name: "Test case firstList = [[1,3],[5,9]], secondList = [] input",
			args: args{
				firstList:  [][]int{{1, 3}, {5, 9}},
				secondList: [][]int{},
			},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.firstList, tt.args.secondList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntersection(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name     string
		args     args
		want     []int
		wantFlag bool
	}{
		{
			name: "Test case first = [12,19], second = [2,7] input",
			args: args{
				first:  []int{12, 19},
				second: []int{2, 7},
			},
			want:     []int{},
			wantFlag: false,
		},
		{
			name: "Test case with first = [12,19], second = [9,14]",
			args: args{
				first:  []int{12, 19},
				second: []int{9, 14},
			},
			want:     []int{12, 14},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [14,16]",
			args: args{
				first:  []int{12, 19},
				second: []int{14, 16},
			},
			want:     []int{14, 16},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [15,24]",
			args: args{
				first:  []int{12, 19},
				second: []int{15, 24},
			},
			want:     []int{15, 19},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [10, 25]",
			args: args{
				first:  []int{12, 19},
				second: []int{10, 25},
			},
			want:     []int{12, 19},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12, 19], second = [10,19]",
			args: args{
				first:  []int{12, 19},
				second: []int{10, 19},
			},
			want:     []int{12, 19},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [12, 40]",
			args: args{
				first:  []int{12, 19},
				second: []int{12, 40},
			},
			want:     []int{12, 19},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [40, 60]",
			args: args{
				first:  []int{12, 19},
				second: []int{40, 60},
			},
			want:     []int{},
			wantFlag: false,
		},
		{
			name: "Test case with first = [12,19], second = [12, 15]",
			args: args{
				first:  []int{12, 19},
				second: []int{12, 15},
			},
			want:     []int{12, 15},
			wantFlag: true,
		},
		{
			name: "Test case with first = [12,19], second = [14,19]",
			args: args{
				first:  []int{12, 19},
				second: []int{14, 19},
			},
			want:     []int{14, 19},
			wantFlag: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFlag := getIntersection(tt.args.first, tt.args.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersection() got = %v, want %v", got, tt.want)
			}
			if gotFlag != tt.wantFlag {
				t.Errorf("getIntersection() gotFlag = %v, want %v", gotFlag, tt.wantFlag)
			}
		})
	}
}
