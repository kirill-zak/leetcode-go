package problem_383

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case with ransomNote = 'a', magazine = 'b'",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "Test case with ransomNote = 'aa', magazine = 'ab'",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "Test case with ransomNote = 'aa', magazine = 'aab'",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
