package problem_12

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case for num = 3 input",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "Test case for num = 58 input",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "Test case for num = 1994 input",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "Test case for num = 2022 input",
			args: args{
				num: 2022,
			},
			want: "MMXXII",
		},
		{
			name: "Test case for num = 702 input",
			args: args{
				num: 702,
			},
			want: "DCCII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createThousand(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case for n = 3000",
			args: args{
				num: 3000,
			},
			want: "MMM",
		},
		{
			name: "Test case for n = 2000",
			args: args{
				num: 2000,
			},
			want: "MM",
		},
		{
			name: "Test case for n = 1000",
			args: args{
				num: 1000,
			},
			want: "M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createThousand(tt.args.num); got != tt.want {
				t.Errorf("createThousand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createHundred(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case for num = 900 input",
			args: args{
				num: 900,
			},
			want: "CM",
		},
		{
			name: "Test case for num = 500 input",
			args: args{
				num: 500,
			},
			want: "D",
		},
		{
			name: "Test case for num = 100 input",
			args: args{
				num: 100,
			},
			want: "C",
		},
		{
			name: "Test case for num = 300 input",
			args: args{
				num: 300,
			},
			want: "CCC",
		},
		{
			name: "Test case for num = 800 input",
			args: args{
				num: 800,
			},
			want: "DCCC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createHundred(tt.args.num); got != tt.want {
				t.Errorf("createHundred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createTen(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case for num = 90",
			args: args{
				num: 90,
			},
			want: "XC",
		},
		{
			name: "Test case for num = 50",
			args: args{
				num: 50,
			},
			want: "L",
		},
		{
			name: "Test case for num = 40",
			args: args{
				num: 40,
			},
			want: "XL",
		},
		{
			name: "Test case for num = 80",
			args: args{
				num: 80,
			},
			want: "LXXX",
		},
		{
			name: "Test case for num = 20",
			args: args{
				num: 20,
			},
			want: "XX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTen(tt.args.num); got != tt.want {
				t.Errorf("createTen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createOne(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case for num = 9",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "Test case for num = 5",
			args: args{
				num: 5,
			},
			want: "V",
		},
		{
			name: "Test case for num = 4",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "Test case for num = 8",
			args: args{
				num: 8,
			},
			want: "VIII",
		},

		{
			name: "Test case for num = 2",
			args: args{
				num: 2,
			},
			want: "II",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createOne(tt.args.num); got != tt.want {
				t.Errorf("createOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
