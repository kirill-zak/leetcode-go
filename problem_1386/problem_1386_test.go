package problem_1386

import "testing"

func Test_maxNumberOfFamilies(t *testing.T) {
	type args struct {
		n             int
		reservedSeats [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]] input",
			args: args{
				n:             3,
				reservedSeats: [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}},
			},
			want: 4,
		},
		{
			name: "Test case with [[2,1],[1,8],[2,6]] input",
			args: args{
				n:             2,
				reservedSeats: [][]int{{2, 1}, {1, 8}, {2, 6}},
			},
			want: 2,
		},
		{
			name: "Test case with [[4,3],[1,4],[4,6],[1,7]] input",
			args: args{
				n:             4,
				reservedSeats: [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}},
			},
			want: 4,
		},
		{
			name: "Test case with [[2.3]] input",
			args: args{
				n:             3,
				reservedSeats: [][]int{{2, 3}},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfFamilies(tt.args.n, tt.args.reservedSeats); got != tt.want {
				t.Errorf("maxNumberOfFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countAvailableVariants(t *testing.T) {
	type args struct {
		bucket Bucket
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case for 0 row",
			args: args{
				bucket: Bucket{
					Row:  0,
					Low:  0,
					High: 11,
				},
			},
			want: 0,
		},
		{
			name: "Test case for [1-10]",
			args: args{
				bucket: Bucket{
					Row:  2,
					Low:  1,
					High: 10,
				},
			},
			want: 2,
		},
		{
			name: "Test case for [0-6]",
			args: args{
				bucket: Bucket{
					Row:  3,
					Low:  0,
					High: 6,
				},
			},
			want: 1,
		},
		{
			name: "Test case for [3-10]",
			args: args{
				bucket: Bucket{
					Row:  2,
					Low:  3,
					High: 10,
				},
			},
			want: 1,
		},
		{
			name: "Test case for [2-9]",
			args: args{
				bucket: Bucket{
					Row:  3,
					Low:  2,
					High: 9,
				},
			},
			want: 1,
		},
		{
			name: "Test case for [5-8]",
			args: args{
				bucket: Bucket{
					Row:  5,
					Low:  5,
					High: 8,
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAvailableVariants(tt.args.bucket); got != tt.want {
				t.Errorf("countAvailableVariants() = %v, want %v", got, tt.want)
			}
		})
	}
}
