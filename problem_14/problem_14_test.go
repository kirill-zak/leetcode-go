package problem_14

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case with strs = ['flower','flow','flight'] input",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Test case with strs = ['dog','racecar','car'] input",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "Test case with strs = ['','b'] input",
			args: args{
				strs: []string{"", "b"},
			},
			want: "",
		},
		{
			name: "Test case with strs = ['b', ''] input",
			args: args{
				strs: []string{"b", ""},
			},
			want: "",
		},
		{
			name: "Test case with strs = ['a'] input",
			args: args{
				strs: []string{"a"},
			},
			want: "a",
		},
		{
			name: "Test case with strs = ['ab', 'a'] input",
			args: args{
				strs: []string{"ab", "a"},
			},
			want: "a",
		},
		{
			name: "Test case with strs = ['flower', 'fkow'] input",
			args: args{
				strs: []string{"flower", "fkow"},
			},
			want: "f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
